package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
)
/**
	生产者

	发送消息示例：下面发送了3条消息，分别发送到了不同的routing_key上面
go run producer.go info this is info message
go run producer.go warning this is warning message
go run producer.go error this is error message
*/
func main() {
	//建立MQ的连接
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)
	defer conn.Close()
	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)
	defer ch.Close()

	//声明一个exchange(交换器)
	err = ch.ExchangeDeclare("routing_exchange","direct",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareExchange)

	//接收消息
	body := utils.BodyForm(os.Args)
	msg := amqp.Publishing{ContentType:"text/plain",Body:[]byte(body)}
	//消息生产者只需要知道把消息发送到指定的exchange即可。
	//exchange和队列的绑定关系在哪里做呢？消息生产者和消息消费者都可以来做，并且在任意一边进行了就行
	err = ch.Publish("routing_exchange",utils.SeverityFrom(os.Args),false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	log.Printf("[x] send %s",body)
}
