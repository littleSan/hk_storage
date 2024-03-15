/*
*

	@author: little
	@date : 2023/9/27
*/
package sysUser

import (
	"hk_storage/common/logger"
	"hk_storage/core/database"
)

func (d *SysUser) GetUserInfoByUid(uid string) (res *SysUser, err error) {
	hsql := "select t.uid ,s.address, s.private_key,t.account from fb_account t left join user_address s on (t.uid = s.uid COLLATE utf8mb4_unicode_ci ) where  t.uid = ?"
	err = database.GetDB().Raw(hsql, uid).First(&res).Error
	if err != nil {
		logger.Info("查询数据出错", err)
		return nil, err
	}
	return
}
