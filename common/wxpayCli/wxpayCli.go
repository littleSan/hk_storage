/*
*

	@author:
	@date : 2023/10/13
*/
package wxpayCli

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"hk_storage/core/configs"
)

func GetClient() *wechat.Client {
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient(configs.TomlConfig.Wxpay.AppAppId, configs.TomlConfig.Wxpay.MchId,
		configs.TomlConfig.Wxpay.MchKey, configs.TomlConfig.Wxpay.IsProd)

	// 打开Debug开关，输出请求日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	client.SetCountry(wechat.China)
	return client
}
