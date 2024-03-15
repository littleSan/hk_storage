/*
*

	@author:
	@date : 2023/10/13
*/
package pay

import (
	"errors"
	"strings"
)

const (
	WXPAY  = "WXPAY"
	ALIPAY = "ALIPAY"
)

var ENUM_PAY_TYPE = map[string]string{
	ALIPAY: ALIPAY,
	WXPAY:  WXPAY,
}

// alipay 请求参数
type AlipayReq struct {
	Oid     string `json:"oid" form:"oid"`
	Subject string `json:"subject" form:"subject"`
	Money   int64  `json:"money" form:"money"` // 分
}

// alipay 响应参数
type AlipayRsp struct {
	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
	PagePayUrl string `json:"page_pay_url"`
	TradeNo    string `json:"trade_no"`
}

type PayRequest struct {
	PayType string `json:"payType" form:"payType"`
	Oid     string `json:"oid" form:"oid"`
	Subject string `json:"subject" form:"subject"`
	Money   int64  `json:"money" form:"money"` // 分
}

// 支付原生币请求舒服
type BaseCoinRequest struct {
	Token   int64  `json:"token" form:"token"`
	Address string `json:"address" form:"address"`
	Uid     string `json:"uid" form:"uid"`
}

type PayResponse struct {
	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
	PagePayUrl string `json:"page_pay_url"`
}

type QueryResponse struct {
	//alipay
	Code          string `json:"code"`
	Msg           string `json:"msg"`
	SubCode       string `json:"sub_code,omitempty"`
	SubMsg        string `json:"sub_msg,omitempty"`
	TradeNo       string `json:"trade_no,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	BuyerLogonId  string `json:"buyer_logon_id,omitempty"`
	TradeStatus   string `json:"trade_status,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty"`
	TransCurrency string `json:"trans_currency,omitempty"`
	//wx
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid       string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId       string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe    string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubOpenid      string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState     string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
}

func (s *PayRequest) ParamCheck() (b1 bool, err error) {
	if ENUM_PAY_TYPE[s.PayType] == "" {
		return false, errors.New("pay type err")
	}
	if strings.Trim(s.Oid, " ") == "" {
		return false, errors.New("oid can not be null")
	}
	return true, nil
}
