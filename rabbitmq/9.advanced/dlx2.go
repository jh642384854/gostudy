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


	//定义一个死信交换器
	err = ch.ExchangeDeclare("dlx_exchange",amqp.ExchangeFanout,true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareExchange)
	//建立一个死信队列
	_,err = ch.QueueDeclare("dlx_queue",true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareQueue)
	//将死信队列和死信交换器绑定
	err = ch.QueueBind("dlx_queue","dlx_exchange_key","dlx_exchange",false,nil)
	utils.FailOnError(err,config.FailExchangeBindQueue)



	//建立一个正常的交换器
	err = ch.ExchangeDeclare("dlx_normal_exchange",amqp.ExchangeFanout,true,false,false,false,nil)
	utils.FailOnError(err,config.FailDeclareExchange)

	//建立一个正常的队列，并为这个队列设置消息过期时间和死信队列
	args := make(amqp.Table)
	args["x-message-ttl"] = 3000
	args["x-dead-letter-exchange"] = "dlx_exchange"
	args["x-dead-letter-routing-key"] = "dlx_exchange_key"
	_,err = ch.QueueDeclare("dlx_normal_queue",true,false,false,false,args)
	utils.FailOnError(err,config.FailDeclareQueue)

	//将正常队列和正常交换器进行绑定
	err = ch.QueueBind("dlx_normal_queue","","dlx_normal_exchange",false,nil)
	utils.FailOnError(err,config.FailExchangeBindQueue)


	//发送消息
	body := []byte("this is a hello demo")
	msg := amqp.Publishing{ContentType:"text/plain",Body:body,DeliveryMode:2}
	err = ch.Publish("dlx_normal_exchange","",false,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	log.Printf("[x] send %s",body)
}