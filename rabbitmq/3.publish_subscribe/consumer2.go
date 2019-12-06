package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
)

/**
	消费者2
	这个声明一个队列(exchange_queue2)，也绑定)到一个公共的exchange上面(publish_exchange)
	两个消费者，声明了两个队列，但是都是绑定在同一个exchange上面，这样在生产消息后，发送到publish_exchange这个exchange上面
	这样绑定在publish_exchange上面的所有队列都会受到消息并进行处理。
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
	queue,err := ch.QueueDeclare("exchange_queue2",true,false,false,false,nil)
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
