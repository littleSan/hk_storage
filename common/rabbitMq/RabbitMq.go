/*
*

	@author:
	@date : 2023/10/20
*/
package rabbitMq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

//type RabbitMq struct {
//	RabbitConn    *amqp.Connection
//	RabbitChannel *amqp.Channel
//}

func Connect() {
	err := errors.New("")
	url := fmt.Sprintf("amqp://%v:%v@%v:%v", configs.TomlConfig.RabbitMq.Username, configs.TomlConfig.RabbitMq.Password,
		configs.TomlConfig.RabbitMq.Host, configs.TomlConfig.RabbitMq.Port)
	RabbitConn, err = amqp.Dial(url)
	if err != nil {
		logger.Error("初始化rabbit 出错", err)
		return
	}
	//初始化 通道
	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel:", err)
	}
	return
}

//// 0表示channel未关闭，1表示channel已关闭
//func CheckRabbitClosed(ch amqp.Channel) int64 {
//	s := reflect.ValueOf(ch)
//	i := s.FieldByName("closed").Int()
//	return i
//}

// GetQueue 声明队列
func GetQueue(queueName string) (queue amqp.Queue, err error) {
	// 声明队列，没有则创建
	// 队列名称、是否持久化、所有消费者与队列断开时是否自动删除队列、是否独享(不同连接的channel能否使用该队列)
	queue, err = RabbitChannel.QueueDeclare(
		queueName, // 队列名字
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Error("get queue err")
		return
	}
	return
}
