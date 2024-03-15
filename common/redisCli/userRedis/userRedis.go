/*
@author: little
@date : 2023/9/27
*/
package userRedis

import (
	"encoding/json"
	"hk_storage/common/constants"
	"hk_storage/common/logger"
	redisCli2 "hk_storage/common/redisCli"
	sysUser2 "hk_storage/models/sysUser"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	SetLoginInfo(key string, user sysUser2.SysUser)
	GetLoginInfo(key string) (user sysUser2.SysUser, err error)
	DeleteKey(key string) bool
	Exits(key string) bool
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
func (s *service) SetLoginInfo(keys string, user sysUser2.SysUser) {
	userByte, _ := json.Marshal(user)
	s.RedisCli.Set(constants.RedisKeyPrefixLoginUser+keys, string(userByte), constants.RedisTimeUint)
}

func (s *service) GetLoginInfo(keys string) (user sysUser2.SysUser, err error) {
	userStr, err := s.RedisCli.Get(constants.RedisKeyPrefixLoginUser + keys)
	if err != nil {
		logger.Info("redis 获取用户缓存出错", err)
		return
	}
	err = json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		logger.Info("解析缓存数据出错", err)
		return
	}
	return
}

func (s *service) Exits(keys string) bool {
	return s.RedisCli.Exists(constants.RedisKeyPrefixLoginUser + keys)
}

func (s *service) DeleteKey(keys string) bool {
	return s.RedisCli.Del(constants.RedisKeyPrefixLoginUser + keys)
}
