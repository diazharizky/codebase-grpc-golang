package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	listenPort := viper.GetInt("listen.port")
	log.Info("Listening on " + viper.GetString("listen.host") + ":" + fmt.Sprintf("%d", listenPort) + "!")
}

func Execute() error {
	return rootCmd.Execute()
}
