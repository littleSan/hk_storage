/*
*

	@author:
	@date : 2023/9/27
*/
package sysUser

import (
	"hk_storage/models/sysUser"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	GetUserInfoByUid(uid string) (res *sysUser.SysUser, err error)
}

type service struct {
	SysUserDao sysUser.SysUserDao
}

func New() *service {
	return &service{
		SysUserDao: sysUser.New(),
	}
}

func (s *service) i() {
}
