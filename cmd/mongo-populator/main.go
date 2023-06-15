package main

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"mongo-oplog-populator/internal/app/populator/generator"
	"mongo-oplog-populator/writer"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var bulkInsert int
var streamInsert int

// TODO: refactor flags add ./mongopop 1000 for bulk , ./mongopop -s 100 for stream
func init() {
	rootCmd.Flags().IntVarP(&streamInsert, "stream", "s", 0, "Stream Insert")
}

var rootCmd = &cobra.Command{
	Use:   "mongopop",
	Short: "Data Population in MongoDB",
	Long:  "This application facilitates data population in a MongoDB database by providing functionalities to perform insert, update, and delete operations. The application allows you to efficiently manage the data in your MongoDB database, resulting in optimized operations and improved performance.",
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().HasFlags() {
			cmd.Usage()
			return
		}

		//TODO-DONE: Write to csv here
		//TODO: Generate data and pass to write
		cfg := config.Load()

		var customGenerator generator.GoFakeItGenerator
		gofakeitGenerator := customGenerator.NewGoFakeItGenerator()
		personnelInfo := gofakeitGenerator.GenerateFakeData()

		csvWriter := writer.NewCSVWriter(cfg.CsvFileName)
		csvWriter.WriteData(personnelInfo)

		ctx, cancel := context.WithCancel(context.Background())
		// Handle interrupt signal
		handleInterruptSignal(cancel)

		//get Client for mongo
		client := writer.GetClient(cfg, ctx)
		//Disconnect Client (pass ctx later for disconnecting mongo client)
		defer writer.DisconnectClient(client, ctx)

		// if csv file does not exist, generate some random/fake data, and populate it to the CSV file

		//TODO: generator will be an interface to genertae data

		//TODO : use reader here
		//TODO : use only 1 flag here
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
