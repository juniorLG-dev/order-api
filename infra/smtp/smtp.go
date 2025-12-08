package smtp

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

type OrderSMTP struct {
	ServerEmail string
	Key         string
}

func NewOrderSMTP(serverEmail, key string) *OrderSMTP {
	return &OrderSMTP{
		ServerEmail: serverEmail,
		Key:         key,
	}
}

func (o *OrderSMTP) SendEmail(
	recipient, subject, message string,
) error {
	m := gomail.NewMessage()
	m.SetHeader("From", o.ServerEmail)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf("<h1>%s<h1>", message))

	d := gomail.NewDialer("smtp.gmail.com", 587, o.ServerEmail, o.Key)
	return d.DialAndSend(m)
}
