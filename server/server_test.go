package server

import (
	"context"
	"testing"

	"github.com/lileio/email_service/email_service"
	"github.com/stretchr/testify/assert"
)

var s = Server{}

func TestRead(t *testing.T) {
	ctx := context.Background()
	req := &email_service.Request{
		Id: "somethingidlike",
	}
	res, err := s.Read(ctx, req)

	assert.Nil(t, err)
	assert.NotEmpty(t, res.Id)
}
