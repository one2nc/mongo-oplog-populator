package cmd

import (
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/populator"
	"os"

	"github.com/spf13/cobra"
)

var operations int

var rootCmd = &cobra.Command{
	Use:   "oplogpop",
	Short: "A simple CLI application",
	Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		//get Client for mongo
		client := config.GetClient()

		// mongoConnection := config.GetMongoConnection(operations, client)

		populator.MakePopulateJson(client, operations)

		//Disconnect Client
		config.DisconnectClient(client)
	},
}

func init() {
	rootCmd.Flags().IntVar(&operations, "op", 0, "No of operations to perform")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
