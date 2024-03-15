package constants

import "time"

const (
	// ProjectName 项目名称
	ProjectName      = "hk_storage"
	HeaderLoginToken = "token"
	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-sysUser:"

	SessionUserInfo = "_session_user_info"
	RedisTimeUint   = time.Hour * 24 * 30
)

// 短信相关
const (
	// SmsRedisTimeUint 短信缓存时长
	SmsRedisTimeUint = time.Minute * 5
	SmsRedisPrefix   = ProjectName + ":sms:"
)

const (
	TaskTimeUint = time.Minute * 30
)
