package plaintext_test

import (
	"net/mail"
	"testing"

	pt "github.com/brentgroves/go_mailtrap/plain_text"
)

// type Email struct {
// 	lastSubject string
// 	lastBody    string
// 	lastTo      []*mail.Address
// }

var subject = "test"
var body = "this is a test"
var emailAddress = []*mail.Address{&mail.Address{Name: "Brent Groves", Address: "gmail.com"}}
var plainTextEmailTest = pt.New(subject, body, emailAddress)

// func (t *Email) Send(subject, body string, to ...*mail.Address)
func TestSendPlainTextEmail(t *testing.T) {

	// email := mt.Email{"test", "this is a test", emailAddress}
	plainTextEmailTest.Send()
	// if sender.lastSubject != "Welcome" {
	// 	t.Error("Subject line was wrong")
	// }
	// if sender.To[0] != to1 && sender.To[1] != to2 {
	// 	t.Error("Wrong recipients")
	// }
}
