package smtp

import (
	"path/filepath"
	"text/template"
)

type ResetPasswordProps struct {
	LOGO string
	URL  string
	NAME string
}

func ResetPasswordTemplate(to []string, props *ResetPasswordProps) error {
	absPath, _ := filepath.Abs("./src/templates/reset-password.html")
	t, err := template.ParseFiles(absPath)
	if err != nil {
		return err
	}

	return SendMail(to, &MailMessage{
		Subject:     "Forgot Password",
		ContentType: HTML,
		HtmlBody:    t,
		HTMLProps:   props,
	})
}
