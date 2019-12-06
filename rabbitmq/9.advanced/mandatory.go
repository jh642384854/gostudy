package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq/config"
	"rabbitmq/utils"
	"time"
)

/**
	https://studygolang.com/articles/12871?fr=sidebar
	mandatory参数：告诉服务器至少将该消息路由到一个队中， 否则将消息返回给生产者。
	immediate参数：告诉服务器如果该消息关联的队列上有消费者，立刻投递:如果所有匹配的队列上都没有消费者，则直接将消息返还给生产者不用将消息存入队列而等待消费者了。
 */

func main() {
	var (
		conn *amqp.Connection
		ch *amqp.Channel
		err error
		notifyConfirm = make(chan amqp.Confirmation)
		notifyReturn = make(chan amqp.Return)
		notifyAck = make(chan uint64)
		notifyNack = make(chan uint64)
		msg amqp.Publishing
	)
	conn,err = amqp.Dial(config.MQ_URL)
	utils.FailOnError(err,config.FailToConnectMQ)

	ch,err = conn.Channel()
	utils.FailOnError(err,config.FailOpenChannel)

	/**
	确认将此通道置于确认模式，以便客户端能够确保服务器已成功接收所有的发布。进入此模式后，服务器将发送一个basic.ack或basic.nack消息，
	其传递标记设置为基于1的增量索引，该索引对应于该方法返回后收到的每个发布。
	向Channel添加侦听器。NotifyPublish响应确认。如果通道。未调用NotifyPublish，确认将被忽略。
	确认订单不受发货订单的约束。Ack和Nack确认将在将来的某个时候到达。不可路由的强制或即时消息在任何通道之后立即得到确认。已通知NotifyReturn监听器。
	当所有应该将消息路由到其上的队列都已收到发送确认，或已将消息加入队列(必要时将消息持久化)时，将确认其他消息。
	当noWait为真时，客户机将不会等待响应。如果服务器不支持此方法，可能会发生通道异常。
	 */
	ch.Confirm(false)

	//定义消息确定被发送
	ch.NotifyPublish(notifyConfirm)
	//当消息无法路由到指定的队列的操作
	ch.NotifyReturn(notifyReturn)
	//定义消息确定被消费
	ch.NotifyConfirm(notifyAck,notifyNack)

	/**
		定义一个备用交换器(就是消息服务达到队列或是到了队列没有消费者消费的情况)
		备份交换器其实和普通的交换器没有太大的区别，为了方便使用，建议设置为 fanout 类型，如若读者想设置为 direct 或者 topic 的类型也没有什么不妥。需要注意的是，消息被重新发送到
		备份交换器时的路由键和从生产者发出的路由键是一样的。
		考虑这样一种情况，如果备份交换器的类型是 direct 并且有一个与其绑定的队列，假设绑定的路由键是 keyl 当某条携带路由键为 key2 的消息被转发到这个备份交换器的时候，备份交换器没有匹配到合适的队列，则消息丢失。如果消息携带的路由键为 key l，则可以存储到队列中。
		对于备份交换器，总结了以下几种特殊情况:
		令如果设置的备份交换器不存在，客户端和 RabbitMQ 服务端都不会有异常出现，此时消息会丢失。
		令如果备份交换器没有绑定任何队列，客户端和 RabbitMQ 服务端都不会有异常出现，此时消息会丢失。
		令如果备份交换器没有任何匹配的队列，客户端和 RabbitMQ 服务端都不会有异常出现，此时消息会丢失。
		如果备份交换器和 mandatory 参数一起使用，那么 mandatory 参数无效。
	 */
	err = ch.ExchangeDeclare("Altemate_Exchange",amqp.ExchangeFanout,true,false,false,false,nil)
	utils.FailOnError(err,config.FialDeclareAltemateExchange)
	//定义一个备用队列
	_,err = ch.QueueDeclare("unroutedqueue",true,false,false,false,nil)
	//将备用队列和备用交换器进行绑定
	err = ch.QueueBind("unroutedqueue","","Altemate_Exchange",false,nil)
	utils.FailOnError(err,"Fail bind Altemate Exchange")


	msg = amqp.Publishing{
		ContentType:"text/plain",
		Body:[]byte("this is a demo"),
	}
	//这个是一个不存在的队列
	err = ch.Publish("","notexitskey",true,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)

	//这个是一个已经存在的队列
	err = ch.Publish("","diyproducer",true,false,msg)
	utils.FailOnError(err,config.FailPublishMessage)


	ticker := time.NewTicker(10*time.Second)
	//for {
		select {
		case ack := <- notifyConfirm:
			if ack.Ack{
				fmt.Println("消息被正常发送")
			}
		case returninfo := <- notifyReturn:
			fmt.Println("消息无法送达，会送达到备用队列里面",string(returninfo.Body),returninfo.RoutingKey,returninfo.Exchange)
			err := ch.Publish("Altemate_Exchange","",false,false,amqp.Publishing{ContentType:returninfo.ContentType,Body:returninfo.Body})
			if err != nil{
				fmt.Println(err)
			}
		case msgid := <- notifyNack:
			fmt.Println("消息未被消费：",msgid)
		case msgid2 := <- notifyAck:
			fmt.Println("消息未被消费：",msgid2)
		case <- ticker.C:
			fmt.Println("每10秒都会执行的内容")
		}
	//}

}
