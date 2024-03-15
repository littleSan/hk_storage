/*
*

	@author:
	@date : 2023/9/27
*/
package sysUser

import (
	"hk_storage/models/sysUser"
)

func (s *service) GetUserInfoByUid(uid string) (res *sysUser.SysUser, err error) {
	return s.SysUserDao.GetUserInfoByUid(uid)
}
