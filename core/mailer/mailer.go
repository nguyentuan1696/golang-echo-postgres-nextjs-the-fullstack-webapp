package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	FromName string
}

type Mailer struct {
	config MailConfig
	auth   smtp.Auth
}

type EmailData struct {
	To           []string
	Subject      string
	Body         string
	Template     string // template file path
	TemplateData any    // data for template
	Attachments  []string
	Cc           []string
	Bcc          []string
}

func NewMailer(config MailConfig) *Mailer {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	return &Mailer{
		config: config,
		auth:   auth,
	}
}

func (m *Mailer) SendMail(data EmailData) error {
	buffer := bytes.NewBuffer(nil)
	boundary := "boundary123"

	// Set headers
	headers := []string{
		fmt.Sprintf("From: %s <%s>", m.config.FromName, m.config.Username),
		fmt.Sprintf("To: %s", strings.Join(data.To, ", ")),
		fmt.Sprintf("Subject: %s", data.Subject),
		"MIME-Version: 1.0",
		fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s", boundary),
		"",
	}

	if len(data.Cc) > 0 {
		headers = append(headers, fmt.Sprintf("Cc: %s", strings.Join(data.Cc, ", ")))
	}

	buffer.WriteString(strings.Join(headers, "\r\n"))

	// Add body
	buffer.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buffer.WriteString("Content-Type: text/html; charset=utf-8\r\n\r\n")
	buffer.WriteString(data.Body)

	// Combine all recipients
	recipients := append(data.To, data.Cc...)
	recipients = append(recipients, data.Bcc...)

	addr := fmt.Sprintf("%s:%d", m.config.Host, m.config.Port)
	return smtp.SendMail(addr, m.auth, m.config.Username, recipients, buffer.Bytes())
}

func (m *Mailer) SendMailWithTemplate(data EmailData) error {
	// Parse template if provided
	if data.Template != "" {
		tmpl, err := template.ParseFiles(data.Template)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		var body bytes.Buffer
		if err := tmpl.Execute(&body, data.TemplateData); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		data.Body = body.String()
	}

	return m.SendMail(data)
}
