package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
	"strings"
)

func main() {
	//建立AMQP协议的连接
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)
	defer conn.Close()

	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)
	defer ch.Close()

	//声明队列
	queue,err := ch.QueueDeclare("workqueue",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareQueue)
	body := bodyForm(os.Args)
	msg := amqp.Publishing{DeliveryMode:amqp.Persistent,ContentType:"text/plain",Body:[]byte(body)}
	err = ch.Publish("",queue.Name,false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)
	log.Printf("[X] send %s",body)
}

func bodyForm(args []string) string {
	var s string
	if (len(args) <2 || os.Args[1] == ""){
		s = "hello"
	}else{
		s = strings.Join(args[1:]," ")
	}
	return s
}