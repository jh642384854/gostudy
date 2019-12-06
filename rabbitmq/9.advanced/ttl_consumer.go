package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
)

func main() {
	//创建MQ连接
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)
	defer conn.Close()

	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)
	defer  ch.Close()

	//获取消息
	msgs,err := ch.Consume("ttl_queue","",false,false,false,false,nil)
	utils.FailOnError(err,config.FailRegisterConsumer)

	forever := make(chan bool)

	go func() {
		for msg := range msgs{
			msg.Ack(true)
			log.Printf("Receive message : %s",msg.Body)
		}
	}()

	log.Printf("[*] Waiting for logs. To exit press Ctil+C")

	<- forever
}
