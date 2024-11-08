package attachments

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	mt "github.com/brentgroves/go_mailtrap"
)

type Attachments mt.Email

func New(subject string, body string, to []*mail.Address) *Attachments {
	at := mt.New(subject, body, to)
	return &Attachments{at.Subject, at.Body, at.To}

}

func (m Attachments) Send() {

	// Open the file
	file, err := os.Open("test.xlsx")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read all contents of the file
	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Encode the file data to base64
	encodedFileData := base64.StdEncoding.EncodeToString(fileData)

	// token := "<secret_token>"
	httpHost := "https://send.api.mailtrap.io/api/send"
	// 	message := []byte(`{"from":{"email":"hello@demomailtrap.com"},
	// "to":[{"email":"brent.groves@gmail.com"}],
	// "subject":"Test",
	// "text":"this is a test"}`)

	message := []byte(`{
    "from": { "email": "hello@demomailtrap.com" },
    "to": [
        { "email": "brent.groves@gmail.com" }
    ],
    "subject": "Attachment Test",
    "text": "Attachment Test.",
    "html": "<p>Check out the attached <strong>file</strong>.</p>",
    "attachments": [
        {
          "filename": "test.xlsx",
          "content": "` + encodedFileData + `",
          "type": "application/xlsx",
          "disposition": "attachment"
        }
    ]
}`)
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
