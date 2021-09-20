package cmd

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Start main app",
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
}

func serve() {
	listenPort := viper.GetString("listen.port")
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svr := grpc.NewServer()
	if err = svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Info("Listening on " + viper.GetString("listen.host") + ":" + listenPort + "!")
}

func Execute() error {
	return rootCmd.Execute()
}
