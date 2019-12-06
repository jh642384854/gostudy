package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmq/config"
	"rabbitmq/utils"
)
/**
	使用示例：
	消费者1：go run consumer.go use.#
	消费者2：go run consumer.go goods.*
 */
func main() {

	//与MQ建立连接
	conn, err := amqp.Dial(config.MQ_URL)
	utils.FailOnError(err, config.FailToConnectMQ)
	defer conn.Close()
	//创建channel
	ch, err := conn.Channel()
	utils.FailOnError(err, config.FailOpenChannel)
	defer ch.Close()

	//声明一个topic交换器
	err = ch.ExchangeDeclare("topic_exchange", amqp.ExchangeTopic, true, false, false, false, nil)
	utils.FailOnError(err, config.FailDeclareExchange)

	//创建队列
	queue, err := ch.QueueDeclare("topic_queue1", true, false, false, false, nil)
	utils.FailOnError(err, config.FailDeclareQueue)

	//将队列和exchange进行绑定
	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}
	for _, routkey := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s", queue.Name, "topic_exchange", routkey)
		err = ch.QueueBind(queue.Name, routkey, "topic_exchange", false, nil)
		utils.FailOnError(err, config.FailExchangeBindQueue)
	}

	//创建消息接收者
	//ch.Get()
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	utils.FailOnError(err, config.FailRegisterConsumer)

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("[*] receiver message:%s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
