package server

import (
	"context"
	"io"
	"testing"

	"github.com/lileio/email_service/email"
	"github.com/lileio/email_service/email_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSender struct {
	mock.Mock
}

func (m MockSender) Close() error {
	return nil
}

func (m MockSender) Send(from string, to []string, msg io.WriterTo) error {
	m.Called()
	return nil
}

var s = Server{}

func TestBasicSend(t *testing.T) {
	mock := MockSender{}
	mock.On("Send").Return(nil)

	email.Open = true
	email.Sender = mock
	go email.Start()

	ctx := context.Background()

	req := &email_service.Email{
		From:      "someone@gmail.com",
		To:        []string{"alex@lile.io"},
		Subject:   "Welcome to lile",
		PlainText: "Hi there, welcome",
	}

	res, err := s.Send(ctx, req)

	mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
