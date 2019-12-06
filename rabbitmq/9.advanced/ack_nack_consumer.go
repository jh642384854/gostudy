package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"rabbitmq/config"
	"rabbitmq/utils"
)

func main() {
	//建立MQ的连接
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)
	defer conn.Close()
	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)
	defer ch.Close()

	msgs ,err := ch.Consume("ack_queue","",false,false,false,false,nil)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		//msg.Ack(false)
		//msg.Nack(false,true)
		msg.Reject(false)
	}

	//
	c := make(chan os.Signal)
	signal.Notify(c,os.Interrupt)
	s := <-c
	log.Println(s)
}
