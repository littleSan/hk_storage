/*
@author: little
@date : 2023/9/27
*/
package sysUser

const (
	StatusUse    = 1
	StatusDelete = 2
	RoleManager  = 1
	RoleCommon   = 2
)

var _ SysUserDao = (*SysUser)(nil)

type SysUser struct {
	Uid     string `gorm:"column:uid" json:"uid,omitempty"`
	Address string `gorm:"column:address" json:"address,omitempty"`    /** 钱包地址*/
	PriKey  string `gorm:"column:private_key" json:"priKey,omitempty"` /** 钱包密钥*/
	Account string `gorm:"column:account" json:"account,omitempty"`    /** 账户*/
}

type SysUserDao interface {
	i()
	GetUserInfoByUid(uid string) (res *SysUser, err error)
}

func New() *SysUser {
	return &SysUser{}
}

func (d *SysUser) i() {
}
