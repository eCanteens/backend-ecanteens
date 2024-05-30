package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type MyEnum int

const (
	HTML MyEnum = iota
	PLAIN
)

func (me MyEnum) String() string {
	return [...]string{
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		"",
	}[me]
}

type MailMessage struct {
	Subject     string
	ContentType MyEnum
	Body        string
	HtmlBody    *template.Template
	HTMLProps   any
}

func SendMail(to []string, message *MailMessage) error {
	AUTH_EMAIL := os.Getenv("SMTP_AUTH_EMAIL")
	AUTH_PASSWORD := os.Getenv("SMTP_AUTH_PASSWORD")
	HOST := os.Getenv("SMTP_HOST")
	PORT := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, HOST)
	smtpAddr := fmt.Sprintf("%s:%s", HOST, PORT)
	subject := fmt.Sprintf("Subject: %s\n", message.Subject)
	mime := message.ContentType.String()

	var body bytes.Buffer

	switch message.ContentType {
	case HTML:
		body.Write([]byte(subject + mime))
		message.HtmlBody.Execute(&body, &message.HTMLProps)
	case PLAIN:
		body.Write([]byte(subject + mime + message.Body))
	}

	err := smtp.SendMail(smtpAddr, auth, AUTH_EMAIL, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
