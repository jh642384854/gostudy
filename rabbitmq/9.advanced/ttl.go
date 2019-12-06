package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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

	//声明一个exchange(交换器),并设定该交换器的消息的TTL值
	args := make(amqp.Table)
	args["x-message-ttl"] = 100000
	err = ch.ExchangeDeclare("ttl_exchange","fanout",true,false,false,false,args)
	utils.FailOnError(err,config.FailDeclareExchange)

	//声明一个队列
	/**
		PRECONDITION_FAILED - inequivalent arg 'x-expires' for queue 'ttl_queue' in vhost '/': received the value '60000' of type 'signedint' but current is none":Failed to declare a queue
	 */
	args2 := make(amqp.Table)
	args2["x-expires"] = 6000
	_,err = ch.QueueDeclare("ttl_queue",true,false,false,false,args2)
	utils.FailOnError(err,config.FailDeclareQueue)

	//将队列和交换器绑定
	err = ch.QueueBind("ttl_queue","","ttl_exchange",false,nil)
	utils.FailOnError(err,config.FailExchangeBindQueue)

	//接收消息
	body := utils.BodyForm(os.Args)
	//注意这里的Expiration的值是一个字符串属性
	msg := amqp.Publishing{ContentType:"text/plain",Body:[]byte(body),Expiration:"5000"}
	//消息生产者只需要知道把消息发送到指定的exchange即可。
	//如果只是将消息发送到exchange中，但是这个exchange并没有和queue相互绑定的话，这条消息会丢失，所以，前面做了exchange和queue的绑定操作
	err = ch.Publish("ttl_exchange","",false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	log.Printf("[x] send %s",body)

}