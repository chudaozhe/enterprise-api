package models

import (
	"crypto/tls"
	"enterprise-api/app/config"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(toMail string, title string, content string) error {
	conf := config.GetConfig().SMTPConfig
	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.Host)

	e := email.NewEmail()
	e.From = conf.Username
	e.To = []string{toMail}
	e.Subject = title
	e.Text = []byte(content)

	if conf.SSL {
		return e.SendWithTLS(conf.Host+conf.SSLPort, auth, &tls.Config{ServerName: conf.Host})
	} else {
		return e.Send(conf.Host+conf.Port, auth)
	}
}
