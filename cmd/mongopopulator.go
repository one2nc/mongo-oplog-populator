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
		mongoConnection := config.GetMongoConnection(args)
		errMakeInsertJson := populator.MakeInsertJson(mongoConnection, args)
		if errMakeInsertJson != nil {
			log.Fatal("err: ", errMakeInsertJson)
		}
		errMakeDeleteJson := populator.MakeDeleteJson(mongoConnection, args)
		if errMakeDeleteJson != nil {
			log.Fatal("err: ", errMakeDeleteJson)
		}
		errMakeInsertAllJson := populator.MakeInsertAllJson(mongoConnection, args)
		if errMakeInsertAllJson != nil {
			log.Fatal("err: ", errMakeInsertAllJson)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
