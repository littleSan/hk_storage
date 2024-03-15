/*
*

	@author:
	@date : 2023/10/10
*/
package pay

import (
	"github.com/gin-gonic/gin"
	"hk_storage/models/pay"
	"hk_storage/services/pay/alipay"
	"hk_storage/services/pay/wxpay"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	PayQrCode(ctx *gin.Context, req *pay.PayRequest) (rest interface{}, err error)
	Notify(resultCode, oid string)
	QueryOrder(ctx *gin.Context, oid string, payType int) (pay.QueryResponse, error)
	CloseOrder(ctx *gin.Context, oid string, payType int) (interface{}, error)
}

type service struct {
	AliPayService alipay.Service
	WxpayService  wxpay.Service
}

func New() *service {
	return &service{
		AliPayService: alipay.New(),
		WxpayService:  wxpay.New()}
}

func (s *service) i() {

}
