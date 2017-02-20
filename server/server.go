package server

import (
	context "golang.org/x/net/context"

	"github.com/go-gomail/gomail"
	"github.com/lileio/email_service/email"
	"github.com/lileio/email_service/email_service"
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

	for k, v := range r.Headers {
		m.SetHeader(k, v)
	}

	resChan := make(chan email.EmailResponse)

	email.EmailQueue <- email.Email{
		Message:      m,
		ResponseChan: resChan,
	}

	res := <-resChan
	if res.Error != nil {
		return nil, res.Error
	}

	return &email_service.EmailResponse{}, nil
}
