/*
*

	@author:
	@date : 2023/10/10
*/
package alipay

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"hk_storage/common/alipayCli"
	"hk_storage/models/pay"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	AlipayTradePreCreate(ctx *gin.Context, req *pay.AlipayReq) (rsp pay.AlipayRsp, err error)
	AlipayPagePayUrl(ctx *gin.Context, req *pay.AlipayReq) (rsp pay.AlipayRsp, err error)
	QueryOrder(ctx *gin.Context, oid string) (*alipay.TradeQueryResponse, error)
	CloseOrder(ctx *gin.Context, oid string) (*alipay.TradeCloseResponse, error)
}

type service struct {
	AliPayCli *alipay.Client
}

func New() *service {
	return &service{AliPayCli: alipayCli.GetClient()}
}

func (s *service) i() {
}
