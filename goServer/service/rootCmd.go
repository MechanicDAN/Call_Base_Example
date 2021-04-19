package service

import (
	"github.com/spf13/cobra"
	"log"
)

var ConfigPath string
var Cmd *cobra.Command

func init() {
	Cmd = &cobra.Command {
		Use: "server",
		Short: "Hello",
		SilenceUsage: true,
	}

	Cmd.PersistentFlags().StringVar(&ConfigPath, "config","", "Config directory")
	err := Cmd.MarkPersistentFlagRequired("config")
	if err != nil {
		log.Fatal(err)
	}

	err = CreateServeCommand()
	if err != nil {
		log.Fatalf("service.rootCmd.init: cant create ServeCommand: %s\n", err)
	}

	err = CreateSchemeCreateCommand()
	if err != nil {
		log.Fatalf("service.rootCmd.init: cant create SchemeCreateCommand %s\n", err)
	}
}



