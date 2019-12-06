package main

import (
	"github.com/streadway/amqp"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	消息生产者
 */
func main() {
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,"Failed to connect Rabbitmq")
	defer conn.Close()

	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,"Failed to open a channel")
	defer ch.Close()

	//声明一个队列
	quene,err := ch.QueueDeclare("hello",true,false,false,false,nil)
	utils.FailOnError(err,"Failed to declare a queue")

	//消息内容
	body := "hello world4"
	msg := amqp.Publishing{
		ContentType:"text/plain",
		Body:[]byte(body),
	}
	err = ch.Publish("",quene.Name,false,false,msg)
	utils.FailOnError(err,"Failed to publish a message")
}

