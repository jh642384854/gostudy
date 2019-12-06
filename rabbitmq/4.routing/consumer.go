package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	消费者
	这个声明一个队列(routing_queue1)，绑定到一个公共的exchange上面(routing_exchange)，并且指定路由到那个key上面。

	示例：
	启动一个消费者：go run consumer.go info error   这个表示这个消费者会处理routing_key为info和error的数据
	启动另外一个消费者：go run consumer.go warning  这个表示这个消费者会处理routing_key为warning的数据

	除此之外的数据都会被丢失。
	后面会介绍如何处理丢失的数据
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
	err = ch.ExchangeDeclare("routing_exchange","direct",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareExchange)
	//创建一个queue
	queue,err := ch.QueueDeclare("routing_queue1",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareQueue)
	if len(os.Args) <2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	//将exchange和queue建立绑定关系
	for _,log_level := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			queue.Name, "logs_direct", log_level)
		err = ch.QueueBind(queue.Name,log_level,"routing_exchange",false,nil)
		utils.FailOnError(err,config.FailExchangeBindQueue)
	}


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
