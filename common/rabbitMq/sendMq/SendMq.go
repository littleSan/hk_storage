/*
*

	@author:
	@date : 2023/10/24
*/
package sendMq

import (
	"github.com/streadway/amqp"
	"hk_storage/common/logger"
	"hk_storage/common/rabbitMq"
)

func SendMsg(queueName, msg string) (err error) {
	//mq := rabbitMq.RabbitMq{}
	//链接mq
	//mq.Connect()
	////关闭mq
	//defer mq.RabbitConn.Close()
	////关闭通道
	//defer mq.RabbitChannel.Close()
	q, err := rabbitMq.GetQueue(queueName)
	if err != nil {
		logger.Error("get queue err")
		return
	}
	//发送消息
	err = rabbitMq.RabbitChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		logger.Error("Failed to publish a message: %v", err)
		return err
	}
	return
}
