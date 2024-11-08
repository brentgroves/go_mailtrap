package common

import (
	"net/mail"
)

type EmailSender interface {
	Send()
	// Send(subject, body string, to ...*mail.Address)
}

type Email struct {
	Subject string
	Body    string
	To      []*mail.Address
}

func New(subject string, body string, to []*mail.Address) *Email {
	return &Email{
		Subject: subject,
		Body:    body,
		To:      to,
	}
}
