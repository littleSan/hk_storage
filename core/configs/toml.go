package configs

import (
	"github.com/BurntSushi/toml"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"path"
)

type TomlConfigs struct {
	HttpAddress string
	Chain       chain
	Db          database `toml:"database"`
	Redis       redis
	Logs        logs
	Language    string
	Active      string
	SmsModel    map[string]smsModel
	SmsContent  map[string]smsContent
	Email       email
	AliOss      aliOss
	Ipfs        ipfs
	Alipay      alipay
	Wxpay       wxpay
	RabbitMq    rabbitMq
}
type chain struct {
	RPCUrl         string
	RPCChainID     int64
	AbiUrl         string
	TransCoinUrl   string
	FileFee        decimal.Decimal
	Profit         decimal.Decimal
	BaseAddress    string
	BaseAddressKey string
}
type database struct {
	Dsn string
}
type redis struct {
	Addr     string
	Password string
	Debug    int
}
type logs struct {
	LogPath  string
	LogLevel string
	OutPut   string
}

type smsModel struct {
	Id       int
	Global   int
	Key      string
	EndPoint string
}
type smsContent struct {
	// 短信头名称，必填
	SignName string
	// 短信内容，需要替换的字符用?code?代替，除阿里云外必填
	Content string
	// 模板编号，阿里云必填
	TemplateCode int
}

type email struct {
	FromAddr string
	Key      string
	Content  string
	Subject  string
}

type aliOss struct {
	Endpoint string
	Id       string
	Secret   string
	Bucket   string
}

type ipfs struct {
	UploadUrl          string
	Authorization      string
	IpfsUrl            string
	PinataGateWay      string
	PinataGateWayToken string
}
type alipay struct {
	Appid              string
	PrivateKey         string
	AppPublicCertPath  string
	AlipayRootCertPath string
	AlipayPublicPath   string
	NotifyUrl          string
	ReturnUrl          string
	SignType           string
	CharSet            string
	IsProd             bool
}
type wxpay struct {
	Appid      string //公众号appId
	PrivateKey string // 公众号appSecret
	MiniAppId  string //小程序appId
	AppAppId   string //app应用appid
	MchId      string //商户号
	MchKey     string //商户密钥
	NotifyUrl  string
	ReturnUrl  string
	IsProd     bool
}

type rabbitMq struct {
	Host     string
	Port     int
	Username string
	Password string
}

var TomlConfig TomlConfigs

// 初始化配置文件
func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := path.Join(pwd, "configs", "config.toml")
	if _, err := toml.DecodeFile(configFile, &TomlConfig); err != nil {
		log.Panic(err)
	}
	//log.Println(TomlConfig)
}
