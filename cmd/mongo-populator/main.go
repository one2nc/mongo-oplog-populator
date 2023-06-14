package main

import (
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"mongo-oplog-populator/writer"
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

		cfg := config.Load()
		//get Client for mongo
		client := writer.GetClient(cfg)
		//Disconnect Client (pass ctx later for disconnecting mongo client)
		defer writer.DisconnectClient(client)

		populator := createPopulator(bulkInsert, streamInsert)

		populator.PopulateData(client, cfg)
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
