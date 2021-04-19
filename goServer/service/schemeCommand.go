package service

import (
	"github.com/spf13/cobra"
	"test/dbDirectory"
	"test/parsers"
)

func CreateSchemeCreateCommand() (err error) {
	var schema = &cobra.Command {
		Use: "schema",
		Short: "create schema",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			parsers.GetConfig(ConfigPath)

			dbDirectory.ConnectToDb()
			defer dbDirectory.DbClose()

			dbDirectory.CreateSchema(args[0])

			return nil
		},
	}

	Cmd.AddCommand(schema)

	return err
}


