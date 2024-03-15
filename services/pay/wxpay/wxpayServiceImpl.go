/*
*

	@author:
	@date : 2023/10/10
*/
package wxpay

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
	pay2 "hk_storage/models/pay"
)

func (s *service) WxGetPaymentQrCode(ctx *gin.Context, req *pay2.PayRequest) (wxRsp *wechat.UnifiedOrderResponse, err error) {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", req.Subject).
		Set("out_trade_no", req.Oid).
		Set("total_fee", req.Money).
		Set("spbill_create_ip", ctx.ClientIP()).
		Set("notify_url", configs.TomlConfig.Wxpay.NotifyUrl).
		Set("trade_type", wechat.TradeType_Native).
		//Set("device_info", "WEB").
		Set("sign_type", wechat.SignType_MD5)
	//SetBodyMap("scene_info", func(bm gopay.BodyMap) {
	//	bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
	//		bm.Set("type", "Wap")
	//		bm.Set("wap_url", "https://www.fmm.ink")
	//		bm.Set("wap_name", "H5测试支付")
	//	})
	//})

	wxRsp, err = s.WxPayCli.UnifiedOrder(ctx, bm)
	if err != nil {
		logger.Info("微信下单失败", err)
		return wxRsp, err
	}
	return wxRsp, err
}

func (s *service) QueryOrder(ctx *gin.Context, oid string) (*wechat.QueryOrderResponse, gopay.BodyMap, error) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("out_trade_no", oid).
		Set("sign_type", wechat.SignType_MD5)
	return s.WxPayCli.QueryOrder(ctx, bm)
}

func (s *service) CloseOrder(ctx *gin.Context, oid string) (*wechat.CloseOrderResponse, error) {
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("out_trade_no", oid).
		Set("sign_type", wechat.SignType_MD5)
	return s.WxPayCli.CloseOrder(ctx, bm)
}
