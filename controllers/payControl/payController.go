/*
*

	@author:
	@date : 2023/10/13
*/
package payControl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat"
	"hk_storage/common/logger"
	"hk_storage/common/rabbitMq/queueConst"
	"hk_storage/common/rabbitMq/sendMq"
	"hk_storage/common/response"
	"hk_storage/core/configs"
	pay2 "hk_storage/models/pay"
	"hk_storage/models/studyFile"
	"hk_storage/services/StudyFile"
	"hk_storage/services/pay"
	"hk_storage/services/sysUser"
	"hk_storage/utils/chainUtil"
	"hk_storage/utils/ethutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	i()
	PayQrCode(ctx *gin.Context)
	WxNotify(ctx *gin.Context)
	QueryOrder(ctx *gin.Context)
	PayBaseCoin(ctx *gin.Context)
}

type controller struct {
	PayService       pay.Service
	StudyFileService StudyFile.Service
	SysUserService   sysUser.Service
}

func New() *controller {
	return &controller{
		PayService:       pay.New(),
		StudyFileService: StudyFile.New(),
		SysUserService:   sysUser.New(),
	}
}

func (c *controller) i() {
}
func (c *controller) PayQrCode(ctx *gin.Context) {
	param := new(pay2.PayRequest)
	err := ctx.BindJSON(param)
	if err != nil {
		logger.Info("二维码支付参数错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	_, err = param.ParamCheck()
	if err != nil {
		logger.Info("二维码支付参数错误", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	rsp, err := c.PayService.PayQrCode(ctx, param)
	if err != nil {
		logger.Info("二维码支付生成二维码信息失败", err)
		ctx.JSON(http.StatusOK, response.FailureMsg(response.Fail, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, rsp))
}
func (c *controller) WxNotify(ctx *gin.Context) {
	bm, err := wechat.ParseNotifyToBodyMap(ctx.Request)
	if err != nil {
		logger.Info("微信支付验签失败", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	logger.Info("支付数据回调", bm)
	// 验签操作
	ok, err := wechat.VerifySign(configs.TomlConfig.Wxpay.MchKey, wechat.SignType_MD5, bm)
	if !ok {
		logger.Info("微信支付验签失败", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	xlog.Debug(bm)
	//key 值 为 NotifyRequest struct 参数
	code := bm.GetString("return_code")
	out_trade_no := bm.GetString("out_trade_no")
	//支付成功
	c.PayService.Notify(code, out_trade_no)
	//sendMq.SendMsg(queueConst.CoinPayQueue, "324593974344159232")
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, response.Success))
}

func (c *controller) AliPayNotify(ctx *gin.Context) {
	//支付宝解析
	notifyReq, err := alipay.ParseNotifyToBodyMap(ctx.Request) // c.Request 是 gin 框架的写法
	if err != nil {
		logger.Info("支付宝支付参数解析失败", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	logger.Info("支付数据回调", notifyReq)
	// 支付宝异步通知验签（公钥模式）
	ok, err := alipay.VerifySign(configs.TomlConfig.Alipay.PrivateKey, notifyReq)
	if ok {
		logger.Info("支付宝验签失败", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	//交易状态
	tradeStatus := notifyReq.Get("trade_status")
	outTradeNo := notifyReq.Get("out_trade_no")
	if tradeStatus == "TRADE_SUCCESS" {
		//付款成功
		c.PayService.Notify("SUCCESS", outTradeNo)
	} else if tradeStatus == "TRADE_CLOSED" {
		//交易超时关闭
		c.PayService.Notify("FAIL", outTradeNo)
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, response.Success))

}

func (c *controller) QueryOrder(ctx *gin.Context) {
	oid, _ := ctx.GetQuery("oid")
	payTypeStr, _ := ctx.GetQuery("payType")
	payType, _ := strconv.Atoi(payTypeStr)

	resP, err := c.PayService.QueryOrder(ctx, oid, payType)
	if err != nil {
		logger.Info("查询失败", err)
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, resP))
}

func (c *controller) CloseOrder(ctx *gin.Context) {
	oid, _ := ctx.GetQuery("oid")
	payTypeStr, _ := ctx.GetQuery("payType")
	payType, _ := strconv.Atoi(payTypeStr)

	resP, err := c.PayService.CloseOrder(ctx, oid, payType)
	if err != nil {
		logger.Info("查询失败", err)
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		return
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, resP))
}

// PayBaseCoin 基础币支付，如yle gls 等，目前针对ylem
func (c *controller) PayBaseCoin(ctx *gin.Context) {
	param := new(pay2.BaseCoinRequest)
	err := ctx.BindJSON(param)
	if err != nil {
		logger.Info("积分支付，参数错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}

	file, err := c.StudyFileService.GetFileByToken(param.Token)
	if file == nil || err != nil {
		logger.Info("积分支付，参数错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	
	if file.Status != studyFile.StatusToChain && file.Status != studyFile.StatusPayFail {
		logger.Info("支付状态错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.PayStatusErr))
		return
	}
	cc, err := c.SysUserService.GetUserInfoByUid(file.Uid)
	if cc == nil || err != nil {
		logger.Info("获取用户信息错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	if strings.Compare(cc.Address, param.Address) != 0 || strings.Compare(file.Address, cc.Address) != 0 {
		logger.Info("用户上链地址不相同")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AddressFormatErr))
		return
	}
	//计算所需的coin 数量

	am1 := ethutil.ToWei(configs.TomlConfig.Chain.FileFee, 18)

	b1 := chainUtil.CheckBalanceEnough(cc.Address, am1.Mul(am1, big.NewInt(int64(file.Amount))))
	if !b1 {
		logger.Info("账户余额不足")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountBalanceNotEnough))
		return
	}

	payHash, err := chainUtil.TransCoin(file.Address, configs.TomlConfig.Chain.BaseAddress, cc.PriKey, am1.Mul(am1, big.NewInt(int64(file.Amount))))
	if err != nil {
		logger.Info("账户余额不足")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.AccountBalanceNotEnough))
		return
	}
	file.PayHash = payHash
	file.Status = studyFile.StatusPaying
	c.StudyFileService.Update(*file)
	//查询支付结果，支付成功，进行上链操作
	sendMq.SendMsg(queueConst.ChainPayQueue, payHash)
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, file))
}
