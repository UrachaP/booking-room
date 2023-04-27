package mailerservice

import (
	"gopkg.in/gomail.v2"
)

type service struct {
	mailer *gomail.Dialer
}

func NewMailerService(mailer *gomail.Dialer) MailerService {
	return &service{mailer: mailer}
}

type MailerService interface {
	Send(message *gomail.Message) error
}

func (m *service) Send(message *gomail.Message) error {
	return m.mailer.DialAndSend(message)
}
