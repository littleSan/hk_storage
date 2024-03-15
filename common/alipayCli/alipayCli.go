/*
*

	@author:
	@date : 2023/10/10
*/
package alipayCli

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"
	"hk_storage/core/configs"
	"os"
)

func GetClient() *alipay.Client {
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	alipayCli, err := alipay.NewClient(configs.TomlConfig.Alipay.Appid, configs.TomlConfig.Alipay.PrivateKey, configs.TomlConfig.Alipay.IsProd)
	if err != nil {
		panic(err)
	}
	// Debug开关，输出/关闭日志
	alipayCli.DebugSwitch = gopay.DebugOn
	// 配置公共参数
	alipayCli.SetCharset(alipay.UTF8).
		SetSignType(alipay.RSA2).
		// SetAppAuthToken("")
		SetReturnUrl(configs.TomlConfig.Alipay.ReturnUrl).
		SetNotifyUrl(configs.TomlConfig.Alipay.NotifyUrl).
		SetLocation(alipay.LocationShanghai)

	// 自动同步验签（只支持证书模式）
	// 传入 支付宝公钥证书 alipayPublicCert.crt 内容
	alipayCli.AutoVerifySign(cert.AlipayPublicContentRSA2)

	// 证书路径
	basePath, _ := os.Getwd()
	err = alipayCli.SetCertSnByPath(basePath+configs.TomlConfig.Alipay.AppPublicCertPath,
		basePath+configs.TomlConfig.Alipay.AlipayRootCertPath, basePath+configs.TomlConfig.Alipay.AlipayPublicPath)
	if err != nil {
		panic(err)
	}
	// 传入证书内容
	//if err = alipayCli.SetCertSnByContent([]byte(configs.TomlConfig.Alipay.AppPublicCertContent),
	//	[]byte(configs.TomlConfig.Alipay.AlipayRootCertContent),
	//	[]byte(configs.TomlConfig.Alipay.AlipayPublicCertContent)); err != nil {
	//	panic(err)
	//}
	return alipayCli
}
