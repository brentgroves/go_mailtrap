package attachments_test

import (
	"net/mail"
	"testing"

	at "github.com/brentgroves/go_mailtrap/attachments"
)

// type Email struct {
// 	lastSubject string
// 	lastBody    string
// 	lastTo      []*mail.Address
// }

var subject = "test"
var body = "this is a test"
var emailAddress = []*mail.Address{{Name: "Brent Groves", Address: "brent.groves@gmail.com"}}
var attachmentEmailTest = at.New(subject, body, emailAddress)

// func (t *Email) Send(subject, body string, to ...*mail.Address)
func TestSendEmailWithAttachment(t *testing.T) {

	// email := mt.Email{"test", "this is a test", emailAddress}
	attachmentEmailTest.Send()
	// if sender.lastSubject != "Welcome" {
	// 	t.Error("Subject line was wrong")
	// }
	// if sender.To[0] != to1 && sender.To[1] != to2 {
	// 	t.Error("Wrong recipients")
	// }
}
