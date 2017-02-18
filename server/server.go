package server

import (
	"errors"

	context "golang.org/x/net/context"

	"github.com/lileio/email_service/email_service"
)

type Server struct {
	email_service.EmailServiceServer
}

var (
	ErrAccountNotFound = errors.New("Something went wrong")
)

func (s Server) Read(ctx context.Context, r *email_service.Request) (*email_service.Response, error) {
	return &email_service.Response{
		Id:	r.Id,
	}, nil
}
