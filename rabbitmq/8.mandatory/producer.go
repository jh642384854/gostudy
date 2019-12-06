package main

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
	"time"
)

/**
	当 mandatory 参数设为 true 时，交换器无法根据自身的类型和路由键找到一个符合条件的队列，那么 RabbitMQ 会调用 Basic.Return 命令将消息返回给生产者。
	当mandatory 数设置为 false 时，出现上述情形，则消息直接被丢弃
	那么生产者如何获取到没有被正确路由到合适队列的消息呢?这时候可以通过调用channel addReturnListener 来添加 ReturnListener 监昕器实现。
	当imrnediate 参数设为 true 时，如果交换器在将消息路由到队列时发现队列上并不存在任何消费者，那么这条消息将不会存入队列中。当与路由键匹配 所有队列都没有消费者时该消息会通过 Basic Return 返回至生产者。

	概括来说 mandatory 参数告诉服务器至少将该消息路由到一个队中，否则将消息返回给生产者。
	imrnediate 参数告诉服务器如果该消息关联的队列上有消费者，立刻投递。如果所有匹配的队列上都没有消费者，则直接将消息返还给生产者 不用将消息存入队列而等待消费者了。
	以下代码参考：https://gist.github.com/OhBonsai/28868448ba84c38749a55ea63f22ca77
 */

type Producer struct {
	name          string
	loger         *log.Logger
	connection    *amqp.Connection
	channel       *amqp.Channel
	done          chan bool
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	isConnected   bool
}

const (
	reconnectDelay = 5 * time.Second //连接断开多久重连
	resendDelay    = 5 * time.Second //消息发生失败后，多久重发
	resendTimes    = 3               //消息重发次数
)

var (
	errNotConnected            = errors.New("Not connected to the producer")
	errAlreadyClosed           = errors.New("already closed:not connected the producer")
	errPushMessageNotConnected = errors.New("Failed to push mesage:not connected")
)

func NewProducer(name, addr string) *Producer {
	producer := Producer{
		loger: log.New(os.Stdout, "", log.LstdFlags),
		name:  name,
		done:  make(chan bool),
	}
	go producer.handleReconnect(addr)
	return &producer
}

//连接rabbitmq。如果连接失败会不断重连，如果连接断开会重新连接
func (producter *Producer) handleReconnect(mqaddr string) {
	for {
		producter.isConnected = false
		log.Println("Attempting to connect")
		for !producter.connect(mqaddr) {
			log.Println("Fail to connect,Retrying...")
			time.Sleep(reconnectDelay)
		}
		select {
		case <-producter.done:
			return
		case <-producter.notifyClose:

		}
	}
}

//连接rabbitmq，以生产者的name定义一个队列
func (producter *Producer) connect(mqaddr string) bool {

	conn, err := amqp.Dial(mqaddr)
	utils.FailOnError(err, config.FailToConnectMQ)

	ch, err := conn.Channel()
	utils.FailOnError(err, config.FailOpenChannel)

	ch.Confirm(false)

	_, err = ch.QueueDeclare(producter.name, true, false, false, false, nil)
	utils.FailOnError(err, config.FailDeclareQueue)

	producter.changeConnection(conn, ch)
	producter.isConnected = true
	log.Println("MQ connected!")
	return true
}

//监听rabbitmq channel的状态
func (producter *Producer) changeConnection(conn *amqp.Connection, ch *amqp.Channel) {
	producter.connection = conn
	producter.channel = ch
	//channel没有必要主动关闭，如果没有协程使用它，它会被垃圾收集器回收
	producter.notifyClose = make(chan *amqp.Error)
	producter.notifyConfirm = make(chan amqp.Confirmation)
	producter.channel.NotifyClose(producter.notifyClose)
	producter.channel.NotifyPublish(producter.notifyConfirm)
}

//三次重发消息
func (producter *Producer) Push(data []byte) error {
	if !producter.isConnected {
		return errPushMessageNotConnected
	}
	var currentTime = 0
	for {
		err := producter.UnsafePush(data)
		if err != nil {
			producter.loger.Println("Push faild,Retrying...")
			currentTime += 1
			if currentTime < resendTimes {
				continue
			} else {
				return err
			}
		}
		ticker := time.NewTicker(resendDelay)
		select {
		case confirm := <-producter.notifyConfirm:
			if confirm.Ack {
				producter.loger.Println("Push Confirmed!")
				return nil
			}
		case <-ticker.C:

		}
		producter.loger.Println("Push didn't confirm. Retrying...")
	}
}

//发送消息，不管是否接收到
func (producter *Producer) UnsafePush(data []byte) error {
	if !producter.isConnected {
		return errNotConnected
	}
	msg := amqp.Publishing{
		DeliveryMode: 2,
		ContentType:  "text/plain", //application/json
		Body:         data,
		Timestamp:    time.Now(),
	}
	return producter.channel.Publish("", producter.name, false, false, msg)
}

//关闭连接/通道
func (producter *Producer) Close() error {
	if !producter.isConnected {
		return errNotConnected
	}
	err := producter.channel.Close()
	if err != nil {
		return err
	}
	err = producter.connection.Close()
	if err != nil {
		return err
	}
	close(producter.done)
	producter.isConnected = false
	return nil
}

func main() {
	producer := NewProducer("diyproducer", config.MQ_URL)
	time.Sleep(2 * time.Second) //因为建立MQ链接是用的协程，如果这里不做休眠处理，下面的代码就不能正常执行，当然可以换更优雅的方式
	err := producer.Push([]byte("hello diy producer"))
	if err != nil {
		log.Println(err)
	}
}
