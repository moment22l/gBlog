package email

import (
	"fmt"
	"gBlog/global"

	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 发送邮箱 接收人,主题,内容
func send(name, subject, body string) error {
	e := global.Conf.Email
	return sendMail(
		e.User,             // 发件人
		e.AuthCode,         // 授权码
		e.Host,             // 服务器地址
		e.Port,             // 服务器端口
		name,               // 接收方
		e.DefaultFromEmail, // 发送人名字
		subject,            // 主题
		body,               // 内容
	)
}

func sendMail(user, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	// 创建一个邮件
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(user, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	fmt.Println(body)
	// 通过gomail的Dialer发送邮件
	d := gomail.NewDialer(host, port, user, authCode)
	err := d.DialAndSend(m)
	return err
}
