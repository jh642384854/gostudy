package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	用法示例：
	发送消息一：go run producer.go use.info "this is user info"

	发送消息二：go run producer.go goods.detail "this is goods detail"

	发送消息三：go run producer.go use.score.list "this is zhangsan score info"
 */

func main() {
	//与MQ建立连接
	conn, err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err, config.FailToConnectMQ)
	defer conn.Close()
	//创建channel
	ch, err := conn.Channel()
	utils.FailOnError(err, config.FailOpenChannel)
	defer ch.Close()

	//声明一个topic交换器
	err = ch.ExchangeDeclare("topic_exchange", amqp.ExchangeTopic, true, false, false, false, nil)
	utils.FailOnError(err, config.FailDeclareExchange)

	//接收消息
	body := utils.BodyForm(os.Args)
	msg := amqp.Publishing{ContentType: "text/plain", Body: []byte(body)}
	//消息生产者只需要知道把消息发送到指定的exchange即可。
	//exchange和队列的绑定关系在哪里做呢？消息生产者和消息消费者都可以来做，并且在任意一边进行了就行
	err = ch.Publish("topic_exchange", utils.SeverityFrom(os.Args), false, false, msg)
	utils.FailOnError(err, config.FailPublishMessage)

	log.Printf("[x] send %s", body)

}
