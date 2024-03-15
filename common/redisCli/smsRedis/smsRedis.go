/*
*

	@author:
	@date : 2023/9/27
*/
package smsRedis

import (
	"hk_storage/common/constants"
	"hk_storage/common/logger"
	redisCli2 "hk_storage/common/redisCli"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	SetSmsCode(key, code string)
	GetSmsCode(key string) (code string, err error)
}

type service struct {
	RedisCli redisCli2.Repo
}

func New() *service {
	return &service{
		RedisCli: redisCli2.New(),
	}
}
func (s *service) i() {
}
func (s *service) SetSmsCode(keys string, code string) {

	s.RedisCli.Set(constants.SmsRedisPrefix+keys, code, constants.SmsRedisTimeUint)
}

func (s *service) GetSmsCode(keys string) (code string, err error) {
	code, err = s.RedisCli.Get(constants.SmsRedisPrefix + keys)
	if err != nil {
		logger.Info("redis 获取用户缓存出错", err)
		return
	}
	return
}
