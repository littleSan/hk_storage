package sms

import (
	"hk_storage/core/configs"
	"strings"
)

// SendMsg
/**
 * @Author little
 * @Description //TODO
 * @Date 15:28 2023/9/26
 * @Param msgType 短信类型 国内 国际
 * @Param phone 手机号
 * @param msg 短信内容
 * @return
 **/
func (i *service) SendMsg(msgType int, active, phone, msg string) bool {
	if active == "" {
		active = configs.TomlConfig.Active
	}
	switch active {
	case "subMail":
		if msgType == 0 {
			return i.SubMail.Send(phone, msg)
		} else {
			return i.SubMail.SendGlobal(phone, msg)
		}
	default:
		return i.SubMail.Send(phone, msg)
	}
}

/**
 * @Author little
 * @Description //TODO
 * @Date 15:50 2023/9/26
 * @Param
 * @return
 **/
func (i *service) SendLoginMsg(msgType int, phone, msg string) bool {
	//拼装 msg信息
	sign := configs.TomlConfig.SmsContent["login"].SignName
	content := configs.TomlConfig.SmsContent["login"].Content
	sign = strings.Split(sign, "/")[msgType]
	//替换拼接content
	content = strings.Split(content, "/")[msgType]
	content = strings.Replace(content, "?code?", msg, 1)
	content = sign + content
	return i.SendMsg(msgType, "", phone, content)
}

func (i *service) SendMail(code, address string) bool {
	return i.EmsService.SendMail(code, address)
}
