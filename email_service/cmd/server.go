package cmd

import (
	"log"

	"github.com/lileio/email_service"
	"github.com/lileio/email_service/email"
	"github.com/lileio/email_service/server"
	"github.com/lileio/lile"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		go email.Start()

		s := &server.Server{}

		impl := func(g *grpc.Server) {
			email_service.RegisterEmailServiceServer(g, s)
		}

		err := lile.NewServer(
			lile.Name("email_service"),
			lile.Implementation(impl),
		).ListenAndServe()

		log.Fatal(err)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
