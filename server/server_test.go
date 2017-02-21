package server

import (
	"context"
	"io"
	"testing"

	"github.com/arbarlow/gomail"
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

type MockDialer struct {
	mock.Mock
	Sender MockSender
}

func (m MockDialer) Dial() (gomail.SendCloser, error) {
	m.Called()
	return m.Sender, nil
}

var s = Server{}

func TestBasicSend(t *testing.T) {
	sender := MockSender{}
	sender.On("Send").Return(nil)

	dialer := MockDialer{Sender: sender}
	dialer.On("Dial").Return(nil, nil)

	email.SetDialerAndSender(&dialer, sender)
	go email.Start()

	ctx := context.Background()

	req := &email_service.Email{
		From:      "someone@gmail.com",
		To:        []string{"alex@lile.io"},
		Subject:   "Welcome to lile",
		PlainText: "Hi there, welcome",
	}

	res, err := s.Send(ctx, req)

	dialer.AssertExpectations(t)
	sender.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestAlternate(t *testing.T)  {}
func TestAttachment(t *testing.T) {}
func TestError(t *testing.T)      {}
