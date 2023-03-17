package cmd

import (
	"log"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/populator"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oplogpop",
	Short: "A simple CLI application",
	Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		//get Client for mongo
		client := config.GetClient()

		mongoConnection := config.GetMongoConnection(args, client)
		err := populator.MakePopulateJson(mongoConnection, args)
		if err != nil {
			log.Fatal("err: ", err)
		}
		//Disconnect Client
		config.DisconnectClient(client)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
