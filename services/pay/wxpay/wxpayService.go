/*
*

	@author:
	@date : 2023/10/10
*/
package wxpay

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"hk_storage/common/wxpayCli"
	pay2 "hk_storage/models/pay"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	WxGetPaymentQrCode(ctx *gin.Context, req *pay2.PayRequest) (rsp *wechat.UnifiedOrderResponse, err error)
	QueryOrder(ctx *gin.Context, oid string) (*wechat.QueryOrderResponse, gopay.BodyMap, error)
	CloseOrder(ctx *gin.Context, oid string) (*wechat.CloseOrderResponse, error)
}

type service struct {
	WxPayCli *wechat.Client
}

func New() *service {
	return &service{WxPayCli: wxpayCli.GetClient()}
}

func (s *service) i() {
}
