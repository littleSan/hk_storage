/*
*

	@author:
	@date : 2023/9/28
*/
package ems

//
//const (
//	mailFromName   = "walkersan@163.com"
//	password       = "FWREPZRCFKDBFLZX"
//	QQmailFromName = "286448087@qq.com"
//	QQpassword     = "ncxvdrjiuwfxbjig"
//
//	mail163 = "walkersan@163.com"
//	mailTqq = "286448087@qq.com"
//)

var _ Service = (*service)(nil)

type Service interface {
	i()
	SendMail(code, address string) bool
	SendMailQQ(code, address string) bool
}

func New() *service {
	return &service{}
}

type service struct {
}

func (s *service) i() {

}
