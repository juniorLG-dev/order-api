package smtp

import (
	"order/application/smtp"
)

type SMTPMocked struct{}

func NewSMTPMocked() smtp.OrderSMTP {
	return &SMTPMocked{}
}

func (s *SMTPMocked) SendEmail(
	recipient, subject, message string,
) error {
	return nil
}
