package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

type MyEnum int

const (
	HTML MyEnum = iota
	Plain
)

func (me MyEnum) String() string {
	return [...]string{"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", ""}[me]
}

type MailMessage struct {
	Subject     string
	ContentType MyEnum
	Body        string
}

func SendMail(to []string, message *MailMessage) error {
	AUTH_EMAIL := os.Getenv("SMTP_AUTH_EMAIL")
	AUTH_PASSWORD := os.Getenv("SMTP_AUTH_PASSWORD")
	HOST := os.Getenv("SMTP_HOST")
	PORT := os.Getenv("SMTP_PORT")

	subject := "Subject: " + message.Subject + "\n"
	mime := message.ContentType.String()
	body := message.Body
	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, HOST)
	smtpAddr := fmt.Sprintf("%s:%s", HOST, PORT)

	err := smtp.SendMail(smtpAddr, auth, AUTH_EMAIL, to, []byte(subject+mime+body))
	if err != nil {
		return err
	}

	return nil
}
