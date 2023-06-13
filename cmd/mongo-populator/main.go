package main

import (
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntVarP(&bulkInsert, "bulk", "b", 0, "Bulk Insert")
	rootCmd.Flags().IntVarP(&streamInsert, "stream", "s", 0, "Stream Insert")
}

var bulkInsert int
var streamInsert int

var rootCmd = &cobra.Command{
	Use:   "mongopop",
	Short: "Populate mongo-oplogs",
	Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().HasFlags() {
			cmd.Usage()
			return
		}

		//get Client for mongo
		client := config.GetClient()

		populator := createPopulator(bulkInsert, streamInsert)

		populator.PopulateData(client)

		//Disconnect Client
		config.DisconnectClient(client)
	},
}

func createPopulator(bulkInsert, streamInsert int) domain.Populator {
	var populator domain.Populator
	if bulkInsert > 0 {
		populator = domain.NewBulkInsert(bulkInsert)
	} else {
		populator = domain.NewStreamInsert(streamInsert)
	}
	return populator
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
