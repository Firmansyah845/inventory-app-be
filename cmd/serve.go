package cmd

import (
	"context"

	"awesomeProjectSamb/cmd/app"
	logger "github.com/sirupsen/logrus"
)

// StartServer : Starts the server both the gRPC and REST servers.
func StartServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	startServer(ctx, cancel)

	<-ctx.Done()
	logger.Infoln("Stopped all server..")
}

func startServer(ctx context.Context, cancel context.CancelFunc) {
	s := app.New()

	s.Start(ctx, cancel)
}
