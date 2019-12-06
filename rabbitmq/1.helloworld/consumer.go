package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	消息消费者
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

	//注册一个消费者
	msgs,err := ch.Consume(quene.Name,"",true,false,false,false,nil)
	utils.FailOnError(err,"Failed to register consumer")
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("Received a message :%s",msg.Body)
		}
	}()
	log.Printf("[*] Waiting for message. To exits press Ctrl+C")

	<- forever
}

