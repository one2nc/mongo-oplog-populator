package main

import (
	"context"
	"fmt"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"mongo-oplog-populator/internal/app/populator/generator"
	"mongo-oplog-populator/writer"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	streamInsert int
	numRecords   int
)

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

		if streamInsert == 0 {
			var err error
			numRecords, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid argument. Please provide a valid integer.")
				os.Exit(1)
			}
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

		populator := createPopulator(numRecords, streamInsert)
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
	if streamInsert > 0 {
		populator = domain.NewStreamInsert(bulkInsert)
	} else {
		populator = domain.NewBulkInsert(streamInsert)
	}
	return populator
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
