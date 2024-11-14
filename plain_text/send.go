package plaintext

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"time"

	mt "github.com/brentgroves/go_mailtrap"
)

type PlainText mt.Email

func New(subject string, body string, to []*mail.Address) *PlainText {
	pt := mt.New(subject, body, to)
	return &PlainText{pt.Subject, pt.Body, pt.To}

}

func (m PlainText) Send() {

	// token := "<secret_token>"
	token := "b789446fb5f5a018106ebee66ce5daec"
	httpHost := "https://send.api.mailtrap.io/api/send"

	// Demo Domain user
	// 	message := []byte(`{"from":{"email":"hello@demomailtrap.com"},
	// "to":[{"email":"brent.groves@gmail.com"}],

	// Repsys.dev Domain User to gmail
	// 	message := []byte(`{"from":{"email":"brent@repsys.dev"},
	// "to":[{"email":"brent.groves@gmail.com"}],
	// "subject":"Test",
	// "text":"this is a test"}`)

	// Repsys.dev Domain User to linamar user
	message := []byte(`{"from":{"email":"brent@repsys.dev"},
"to":[{"email":"brent.groves@linamar.com"}],
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
