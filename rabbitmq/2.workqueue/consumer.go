package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/config"
	"rabbitmq/utils"
	"time"
)

func main() {
	//与MQ建立连接
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

	//设置QOS。即设置流量控制
	err = ch.Qos(1,0,false)
	utils.FailOnError(err,config.FailSetQos)

	//创建消费者
	msgs,err := ch.Consume(queue.Name,"",false,false,false,false,nil)
	utils.FailOnError(err,config.FailRegisterConsumer)

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("Receive a message : %s",msg.Body)
			//获取消息后面的.(点号多少)来决定sleep多长时间
			dot_count := bytes.Count(msg.Body,[]byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t*time.Second)
			log.Printf("Work Done")
			msg.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for message.To exit press Ctrl+C")

	<- forever
}
