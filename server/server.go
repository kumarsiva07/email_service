package server

import (
	"bytes"

	context "golang.org/x/net/context"

	"github.com/arbarlow/gomail"
	"github.com/lileio/email_service"
	"github.com/lileio/email_service/email"
)

type Server struct {
	email_service.EmailServiceServer
}

func (s Server) Send(ctx context.Context, r *email_service.Email) (*email_service.EmailResponse, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", r.From)
	m.SetHeader("To", r.To...)
	m.SetHeader("Subject", r.Subject)
	m.SetBody("text/plain", r.PlainText)

	if r.HtmlAlternate != "" {
		m.AddAlternative("text/html", r.HtmlAlternate)
	}

	for _, a := range r.Attachments {
		m.AttachReader(a.Filename, bytes.NewReader(a.Body))
	}

	for k, v := range r.Headers {
		m.SetHeader(k, v)
	}

	err := email.SendMessage(m)
	if err != nil {
		return nil, err
	}

	return &email_service.EmailResponse{}, nil
}
