package response

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"hk_storage/core/configs"
)

type Data struct {
	Code    int         `json:"code"`    // 业务码
	Message string      `json:"message"` // 描述信息
	Data    interface{} `json:"data"`    // 描述信息
}

func SUCCESS(ctx *gin.Context, val interface{}) Data {
	language := ctx.GetHeader("language")

	return Data{
		Code:    Success,
		Message: Text(language, Success),
		Data:    val,
	}
}

func Failure(ctx *gin.Context, code int) Data {
	language := ctx.GetHeader("language")
	return Data{
		Code:    code,
		Message: Text(language, code),
	}
}

func FailureMsg(code int, msg string) Data {
	return Data{
		Code:    code,
		Message: msg,
		Data:    "",
	}
}

const (
	Success            = 200200
	Fail               = 400400
	ParamError         = 400402
	AuthorizationError = 400401
	AccountError       = 400403
	NotRegister        = 400404
	PasswordErr        = 400405
	SmsCodeErr         = 400406
	UsernameErr        = 400407
	AccountRegistered  = 400408
	AccountFormatErr   = 400409
	UpdateErr          = 400410
	FileFormatErr      = 400411

	ServerNotExist          = 400412
	BenefitCodeErr          = 400413
	BenefitCodeUsed         = 400414
	BenefitNotEnough        = 400415
	PermissionErr           = 400416
	AddressFormatErr        = 400417
	AccountBalanceNotEnough = 400418
	PayStatusErr            = 400419
	FileAmountErr           = 400420
	FileNotExist            = 400421
	FileStatusErr           = 400422
)

func Text(language string, code int) string {
	if language == "" {
		language = configs.TomlConfig.Language
	}
	if language == "zh-cn" {
		return zhCNText[code]
	}

	if language == "en" {
		return enUSText[code]
	}

	return zhCNText[code]
}
