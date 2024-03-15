/*
*

	@author:
	@date : 2023/10/24
*/
package receiver

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"hk_storage/common/client"
	"hk_storage/common/logger"
	"hk_storage/common/rabbitMq"
	"hk_storage/common/rabbitMq/queueConst"
	"hk_storage/common/rabbitMq/sendMq"
	"hk_storage/models/studyFile"
	"hk_storage/services/StudyFile"
	"strings"
	"time"
)

var StudyService = StudyFile.New()

// 接受数据
func ReceiveMsg(queueName string) {
	//mq := new(rabbitMq.RabbitMq)
	//defer rabbitMq.RabbitConn.Close()
	//defer rabbitMq.RabbitChannel.Close()
	q, err := rabbitMq.GetQueue(queueName)
	if err != nil {
		logger.Error("get queue err")
		return
	}

	// 队列名称、consumer、auto-ack、是否独享
	// deliveries是一个管道，有消息到队列，就会消费，消费者的消息只需要从deliveries这个管道获取
	deliveries, err := rabbitMq.RabbitChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack,true消费了就消失
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)
	go func() {
		for d := range deliveries {
			logger.Info(fmt.Sprintf("返回的消息:%s", d.Body))
			msgData := string(d.Body)
			if strings.Trim(msgData, " ") != "" {
				logger.Info("接受上链结果数据，", msgData)
				ChainConfirm(msgData)
			}
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ChainConfirm(hash string) {
	res, err := StudyService.GetFileByHash(hash)
	if err != nil {
		logger.Info("处理上链数据，根据hash 未查询到", err)
		return
	}
	// 查询交易是否成功，成功更改状态，失败重新发送，
	tx, err := client.GetEthClient().TransactionReceipt(context.Background(), common.HexToHash(res.Hash))
	if err != nil {
		logger.Info("暂未获取到链交易结果，进入等待", res.Hash)
		time.Sleep(1 * time.Second)
		sendMq.SendMsg(queueConst.ChainQueue, hash)
		return
	}
	if tx.Status == 1 {
		processSuccess(hash)
	} else {
		processFail(hash)
	}
}

func processSuccess(hash string) {
	res, _ := StudyService.GetFileByHash(hash)
	res.Status = studyFile.StatusChainSuccess
	StudyService.Update(*res)
}

func processFail(hash string) {
	res, _ := StudyService.GetFileByHash(hash)
	res.Status = studyFile.StatusChainFail
	StudyService.Update(*res)
}
