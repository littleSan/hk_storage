/*
*

	@author:
	@date : 2023/10/27
*/
package consumer

import (
	"hk_storage/common/rabbitMq/queueConst"
	"hk_storage/common/rabbitMq/receiver"
)

// 消费者初始化
func ConsumersInit() {
	go receiver.ReceiveMsg(queueConst.ChainQueue)
	go receiver.ReceivePayMsg(queueConst.ChainPayQueue)
}
