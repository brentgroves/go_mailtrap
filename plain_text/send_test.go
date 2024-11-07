package plaintext_test

import (
	"net/mail"
	"testing"

	mt "github.com/brentgroves/go_mailtrap/plain_text"
)

// type Email struct {
// 	lastSubject string
// 	lastBody    string
// 	lastTo      []*mail.Address
// }

var emailAddress = []*mail.Address{&mail.Address{Name: "Brent Groves", Address: "gmail.com"}}
var emailTests = mt.PlainTextEmail{
	Subject: "test",
	Body:    "this is a test",
	To:      emailAddress,
}

// func (t *Email) Send(subject, body string, to ...*mail.Address)
func TestSendEmail(t *testing.T) {

	// email := mt.Email{"test", "this is a test", emailAddress}
	emailTests.Send()
	// if sender.lastSubject != "Welcome" {
	// 	t.Error("Subject line was wrong")
	// }
	// if sender.To[0] != to1 && sender.To[1] != to2 {
	// 	t.Error("Wrong recipients")
	// }
}
