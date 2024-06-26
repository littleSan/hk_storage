package response

var zhCNText = map[int]string{
	Success:                 "成功",
	Fail:                    "失败",
	ParamError:              "参数信息错误",
	AuthorizationError:      "登陆token失效",
	AccountError:            "账户有误",
	NotRegister:             "暂未注册",
	PasswordErr:             "密码错误",
	SmsCodeErr:              "短信验证码错误",
	UsernameErr:             "用户名错误",
	AccountRegistered:       "账户已注册",
	AccountFormatErr:        "账户格式错误",
	UpdateErr:               "更新失败",
	FileFormatErr:           "文件格式错误或大小超过限制",
	ServerNotExist:          "服务包不存在",
	BenefitCodeErr:          "权益码错误",
	BenefitCodeUsed:         "权益码已使用",
	BenefitNotEnough:        "权益次数不足",
	PermissionErr:           "权限不足",
	AddressFormatErr:        "地址错误格式",
	AccountBalanceNotEnough: "账户可用余额不足",
	PayStatusErr:            "支付状态错误",
	FileAmountErr:           "文件数量错误",
	FileNotExist:            "文件不存在",
	FileStatusErr:           "文件状态值错误",
	YLemReceived:            "今日已领取",
	YLemMaxLimitErr:         "今日领取已达上限",
}
