package sms

import (
	"hk_storage/services/sms/ems"
	"hk_storage/services/sms/subMail"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	SendMsg(msgType int, active, phone, msg string) bool
	SendLoginMsg(msgType int, phone, msg string) bool
	SendMail(code, address string) bool
}

type service struct {
	SubMail    subMail.Service
	EmsService ems.Service
}

func New() *service {
	return &service{
		SubMail:    subMail.New(),
		EmsService: ems.New(),
	}
}

func (i *service) i() {

}
