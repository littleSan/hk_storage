/*
*

	@author:
	@date : 2023/10/10
*/
package alipay

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/shopspring/decimal"
	"hk_storage/models/pay"
)

// AlipayGetPaymentQrCode 商家当面付
func (s *service) AlipayTradePreCreate(ctx *gin.Context, req *pay.AlipayReq) (rsp pay.AlipayRsp, err error) {
	// 生成单号
	tradeNo := req.Oid
	//处理金额
	amount := decimal.NewFromInt(req.Money).DivRound(decimal.NewFromInt(100), 2).String()
	xlog.Infof("tradeNo: %s, amount: %s", tradeNo, amount)
	// 构造参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", req.Subject).
		Set("out_trade_no", tradeNo).
		Set("total_amount", amount)
	// 发起支付
	aliRsp, err := s.AliPayCli.TradePrecreate(ctx, bm)
	if err != nil {
		if bizError, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("s.alipay.TradePrecreate(%v), bizError:%v", bm, bizError)
			return rsp, err
		}
		xlog.Errorf("s.alipay.TradePrecreate(%v), err:%v", bm, err)
		return rsp, err
	}
	// return
	rsp = pay.AlipayRsp{
		OutTradeNo: aliRsp.Response.OutTradeNo,
		QrCode:     aliRsp.Response.QrCode,
	}
	return rsp, nil
}

// AlipayPagePayUrl 获取支付宝网页支付链接
func (s *service) AlipayPagePayUrl(ctx *gin.Context, req *pay.AlipayReq) (rsp pay.AlipayRsp, err error) {
	// 生成单号
	tradeNo := req.Oid
	amount := decimal.NewFromInt(req.Money).DivRound(decimal.NewFromInt(100), 2).String()
	xlog.Infof("tradeNo: %s, amount: %s", tradeNo, amount)
	// 构造参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", req.Subject).
		Set("out_trade_no", tradeNo).
		Set("total_amount", amount)
	// 发起支付
	pagePayUrl, err := s.AliPayCli.TradePagePay(ctx, bm)
	if err != nil {
		xlog.Errorf("s.alipay.TradePagePay(%v), err:%v", bm, err)
		return rsp, err
	}
	// return
	rsp = pay.AlipayRsp{
		OutTradeNo: tradeNo,
		PagePayUrl: pagePayUrl,
	}
	return rsp, nil
}

func (s *service) QueryOrder(ctx *gin.Context, oid string) (*alipay.TradeQueryResponse, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", oid)
	return s.AliPayCli.TradeQuery(ctx, bm)
}

func (s *service) CloseOrder(ctx *gin.Context, oid string) (*alipay.TradeCloseResponse, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", oid)
	return s.AliPayCli.TradeClose(ctx, bm)
}
