package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	消费者1
	这个声明一个队列(exchange_queue1)，绑定到一个公共的exchange上面(publish_queue)
 */

func main() {
	//创建MQ连接
	conn,err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)
	defer conn.Close()

	//创建channel
	ch,err := conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)
	defer  ch.Close()

	//声明一个exchange
	err = ch.ExchangeDeclare("publish_exchange","fanout",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareExchange)
	//创建一个queue
	queue,err := ch.QueueDeclare("exchange_queue1",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareQueue)
	//将exchange和queue建立绑定关系
	err = ch.QueueBind(queue.Name,"","publish_exchange",false,nil)
	utils.FailOnError(err,config.FailExchangeBindQueue)

	//获取消息
	msgs,err := ch.Consume(queue.Name,"",false,false,false,false,nil)
	utils.FailOnError(err,config.FailRegisterConsumer)

	forever := make(chan bool)

	go func() {
		for msg := range msgs{
			log.Printf("Receive message : %s",msg.Body)
		}
	}()

	log.Printf("[*] Waiting for logs. To exit press Ctil+C")

	<- forever
}
