package main

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"mongo-oplog-populator/writer"
	"os"
	"os/signal"
	"syscall"

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

		ctx, cancel := context.WithCancel(context.Background())
		// Handle interrupt signal
		handleInterruptSignal(cancel)

		cfg := config.Load()
		//get Client for mongo
		client := writer.GetClient(cfg, ctx)
		//Disconnect Client (pass ctx later for disconnecting mongo client)
		defer writer.DisconnectClient(client, ctx)

		// if csv file does not exist, generate some random/fake data, and populate it to the CSV file
		noOfOperations := getNumberOfOperations(bulkInsert, streamInsert)
		csvWriter := writer.NewCSVWriter(cfg.CsvFileName, noOfOperations)
		csvWriter.WriteData()

		populator := createPopulator(bulkInsert, streamInsert)
		populator.PopulateData(client, cfg, ctx)
	},
}

func handleInterruptSignal(cancel context.CancelFunc) {
	// Create an interrupt channel to listen for the interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		// Cancel the context to signal the shutdown
		cancel()
	}()
}

func getNumberOfOperations(bulkInsert, streamInsert int) int {
	if bulkInsert > 0 {
		return bulkInsert
	} else {
		return streamInsert
	}

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
