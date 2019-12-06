package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq/config"
	"rabbitmq/utils"
	"time"
)

/**
	保证消息准确收到的方式二：发送方确认机制
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

	ch.Confirm(false)

	confirm := make(chan amqp.Confirmation)
	ch.NotifyPublish(confirm)

/*	ch.ExchangeDeclare("ack_exchange",amqp.ExchangeFanout,true,false,false,false,nil)
	ch.QueueDeclare("ack_queue",true,false,false,false,nil)
	ch.QueueBind("ack_queue","","ack_exchange",false,nil)
*/
	body := []byte("this is a ack message")
	msg := amqp.Publishing{ContentType:"text/plain",DeliveryMode:2,Body:body}
	err = ch.Publish("ack_exchange","",false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	ticker := time.Tick(1*time.Second)
	//forever := make(chan bool)
	select {
	case confirminfo := <- confirm:
		fmt.Println(confirminfo.Ack,confirminfo.DeliveryTag)
	case <- ticker:
		fmt.Println("every time second execute")
	}
	//阻塞
	//<- forever
}