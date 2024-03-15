/*
*

	@author:
	@date : 2024/3/12
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
	"hk_storage/models/contract"
	"hk_storage/models/studyFile"
	"hk_storage/utils/chainUtil"
	"strings"
	"time"
)

// 接受数据
func ReceivePayMsg(queueName string) {
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
				ChainPayConfirm(msgData)
			}
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ChainPayConfirm(hash string) {
	res, err := StudyService.GetFileByPayHash(hash)
	if err != nil {
		logger.Info("处理上链数据，根据hash 未查询到", err)
		return
	}
	// 查询交易是否成功，成功更改状态，失败重新发送，
	tx, err := client.GetEthClient().TransactionReceipt(context.Background(), common.HexToHash(res.PayHash))
	if err != nil {
		logger.Info("暂未获取到链交易结果，进入等待", res.PayHash)
		time.Sleep(1 * time.Second)
		sendMq.SendMsg(queueConst.ChainPayQueue, hash)
		return
	}
	if tx.Status == 1 {
		//付款成功，进行上链操作
		PaySuccess(hash)
	} else {
		//付款失败，不进行上链操作
		PayFail(hash)
	}

}

func PaySuccess(hash string) {
	res, _ := StudyService.GetFileByPayHash(hash)
	res.Status = studyFile.StatusPaySuccess
	StudyService.Update(*res)
	//付款成功，进行上链操作
	c1 := &contract.Contract{}
	c1.Type = "storage"
	c1.GetInfo()
	hash, err := chainUtil.Mint(c1, "mint", float64(res.Token), res.JsonUrl, c1.From, c1.PriKey)
	if err != nil {
		logger.Info("上链操作失败", err)
		return
	}
	res.Status = studyFile.StatusChaining
	res.ContractAddr = c1.Address
	res.Hash = hash
	StudyService.Update(*res)
	sendMq.SendMsg(queueConst.ChainQueue, hash)
}

func PayFail(hash string) {
	res, _ := StudyService.GetFileByPayHash(hash)
	res.Status = studyFile.StatusPayFail
	StudyService.Update(*res)
}
