package service

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"test/dbDirectory"
	"test/parsers"
	"test/server"
)

func CreateServeCommand() (err error) {
	var serve = &cobra.Command {
		Use: "serve",
		Short: "start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			parsers.GetConfig(ConfigPath)

			dbDirectory.ConnectToDb()
			defer dbDirectory.DbClose()

			fmt.Println("run at " + parsers.GetHost() + parsers.GetPort())
			err = server.StartServer()
			if err != nil {
				log.Printf("service.serviceCommand.CreateServeCommand: +%v\n",err)
				return err
			}
			return nil
		},
	}

	Cmd.AddCommand(serve)

	return err
}
