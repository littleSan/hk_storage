/*
*

	@author:
	@date : 2023/10/10
*/
package pay

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"hk_storage/common/logger"
	"hk_storage/models/pay"
)

func (s *service) PayQrCode(ctx *gin.Context, req *pay.PayRequest) (rsp interface{}, err error) {
	//支付宝二维码支付

	if req.PayType == pay.ALIPAY {

		tem := &pay.AlipayReq{Oid: req.Oid, Subject: req.Subject, Money: req.Money}
		rsp, err = s.AliPayService.AlipayPagePayUrl(ctx, tem)
		if err != nil {
			logger.Info("获取支付宝支付二维码错误", err)
			return rsp, err
		}
		//return rsp, err
	} else if req.PayType == pay.WXPAY {

		rsp, err = s.WxpayService.WxGetPaymentQrCode(ctx, req)
		if err != nil {
			logger.Info("获取支付宝支付二维码错误", err)
			return rsp, err
		}
		logger.Info("微信支付结果", rsp)

	}
	//更新用户订单状态
	return rsp, nil
}

func (s *service) Notify(resultCode, oid string) {
	//传入hash信息，通过rabbit 进行操作
	//sendMq.SendMsg(queueConst.ChainQueue, oid)
}

func (s *service) QueryOrder(ctx *gin.Context, oid string, payType int) (res pay.QueryResponse, err error) {
	var a interface{}
	if payType == 1 {
		temp, _ := s.AliPayService.QueryOrder(ctx, oid)
		a = temp.Response
	} else if payType == 2 {
		a, _, err = s.WxpayService.QueryOrder(ctx, oid)
	} else {
		return res, errors.New("payType err")
	}
	byteDate, err := json.Marshal(a)
	if err != nil {
		logger.Info("解析支付响应数据出错", err)
		return res, nil
	}
	err = json.Unmarshal(byteDate, &res)
	if err != nil {
		logger.Info("解析支付响应数据出错", err)
		return res, nil
	}
	return res, nil
}

func (s *service) CloseOrder(ctx *gin.Context, oid string, payType int) (resP interface{}, err error) {
	if payType == 1 {
		resP, err = s.AliPayService.CloseOrder(ctx, oid)
	} else if payType == 2 {
		resP, err = s.WxpayService.CloseOrder(ctx, oid)
	} else {
		return nil, errors.New("payType err")
	}
	return
}
