package main

import (
	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
	"github.com/lileio/email_service/email"
	"github.com/lileio/email_service/email_service"
	"github.com/lileio/email_service/server"
	"github.com/lileio/lile"
)

func main() {
	go email.Start()
	log.SetLevel(log.DebugLevel)

	s := &server.Server{}

	impl := func(g *grpc.Server) {
		email_service.RegisterEmailServiceServer(g, s)
	}

	err := lile.NewServer(
		lile.Name("email_service"),
		lile.Implementation(impl),
	).ListenAndServe()

	log.Fatal(err)
}
