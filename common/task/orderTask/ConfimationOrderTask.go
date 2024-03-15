/*
*

	@author:
	@date : 2023/11/9

@desc : 确权订单回退
*/
package orderTask

import (
	"time"
)

func NewCronJob(duration time.Duration, job func()) (stopChan chan bool) {
	ticker := time.NewTicker(duration)
	stopChan = make(chan bool)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		// 由于ticker.Stop()内部不会关闭chan，故使用 for range 会内存泄露
		// 推荐使用 for + select + return 的方式，让 ticker 最终被GC
		for {
			select {
			case <-ticker.C:
				job()
			case stop := <-stopChan:
				if stop {
					close(stopChan)
					return
				}
			}
		}
	}(ticker)
	return stopChan
}

func ConfirmationOrderBack() {
	//
}
