package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	死信队列
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

	ch.ExchangeDeclare("exchange.dlx",amqp.ExchangeDirect,true,false,false,false,nil)

	ch.ExchangeDeclare("exchange.normal",amqp.ExchangeFanout,true,false,false,false,nil)

	args := make(amqp.Table)
	args["x-message-ttl"] = 10000
	args["x-dead-letter-exchange"] = "exchange.dlx"
	args["x-dead-letter-routing-key"] = "routingkey"

	ch.QueueDeclare("queue.normal",true,false,false,false,args)
	ch.QueueBind("queue.normal","","exchange.normal",false,nil)

	ch.QueueDeclare("queue.dlx",true,false,false,false,nil)
	ch.QueueBind("queue.dlx","routingkey","exchange.dlx",false,nil)

	//发送消息
	body := []byte("this is a hello demo")
	msg := amqp.Publishing{ContentType:"text/plain",Body:body,DeliveryMode:2}
	err = ch.Publish("exchange.normal","",false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	log.Printf("[x] send %s",body)
}