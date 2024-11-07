package plaintext

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"time"
)

// type EmailSender interface {
// 	Send()
// 	// Send(subject, body string, to ...*mail.Address)
// }

type PlainTextEmail struct {
	Subject string
	Body    string
	To      []*mail.Address
}

func (m PlainTextEmail) Send() {

	// token := "<secret_token>"
	httpHost := "https://send.api.mailtrap.io/api/send"
	message := []byte(`{"from":{"email":"hello@demomailtrap.com"},
"to":[{"email":"brent.groves@gmail.com"}],
"subject":"Test",
"text":"this is a test"}`)

	// Set up request
	request, err := http.NewRequest(http.MethodPost, httpHost, bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}
	// Set required headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
	// Send request
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
