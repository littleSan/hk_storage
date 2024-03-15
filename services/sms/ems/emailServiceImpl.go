/*
*

	@author:
	@date : 2023/9/28
*/
package ems

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"hk_storage/core/configs"
	"net/smtp"
	"strings"
)

func (s *service) SendMail(code, address string) bool {
	e := email.NewEmail()

	mailUserName := configs.TomlConfig.Email.FromAddr //邮箱账号
	mailPassword := configs.TomlConfig.Email.Key      //邮箱授权码
	e.From = mailUserName
	e.To = []string{address}
	e.Subject = configs.TomlConfig.Email.Subject
	content := configs.TomlConfig.Email.Content
	content = strings.Replace(content, "?code?", code, 1)
	e.HTML = []byte(content)
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	return true
}

func (s *service) SendMailQQ(code, address string) bool {

	mailUserName := configs.TomlConfig.Email.FromAddr //邮箱账号
	mailPassword := configs.TomlConfig.Email.Key      //邮箱授权码
	addr := "smtp.qq.com:465"                         //TLS地址
	host := "smtp.qq.com"                             //邮件服务器地址
	e := email.NewEmail()
	e.From = mailUserName
	e.To = []string{address}
	e.Subject = configs.TomlConfig.Email.Subject
	content := configs.TomlConfig.Email.Content
	content = strings.Replace(content, "?code?", code, 1)
	e.HTML = []byte(content)
	err := e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	return true
}
