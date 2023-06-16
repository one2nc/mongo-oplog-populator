package main

import (
	"context"
	"fmt"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/domain"
	"mongo-oplog-populator/internal/app/populator/generator"
	"mongo-oplog-populator/internal/app/populator/reader"
	"mongo-oplog-populator/internal/app/populator/service"
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

// TODO-DONE: refactor flags add ./mongopop 1000 for bulk , ./mongopop -s 100 for stream
func init() {
	//TODO-DONE : use only 1 flag here
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

		cfg := config.Load()

		//TODO-DONE: Generate data and pass to write
		var customGenerator generator.GoFakeItGenerator
		gofakeitGenerator := customGenerator.NewGoFakeItGenerator()
		personnelInfo := gofakeitGenerator.GenerateFakeData()

		//TODO-DONE: Write to csv here
		//checking if csv file already exists, if not, ythen create a writer
		_, err := os.Stat(cfg.CsvFileName)
		if os.IsNotExist(err) {
			csvWriter := writer.NewCSVWriter(cfg.CsvFileName)
			csvWriter.WriteData(personnelInfo)
		}

		//TODO-DONE: use reader here
		csvReader := reader.NewCSVReader(cfg.CsvFileName)
		personnelInfo = csvReader.ReadData()

		ctx, cancel := context.WithCancel(context.Background())
		// Handle interrupt signal
		handleInterruptSignal(cancel)

		//get Client for mongo
		client := writer.GetClient(ctx, cfg)

		//Disconnect Client (pass ctx later for disconnecting mongo client)
		defer writer.DisconnectClient(ctx, client)

		//TODO-DONE: remove hardcoded number from here
		noOfOperations := getNoOfOperations(streamInsert, numRecords)
		opSize := service.CalculateOperationSize(noOfOperations)

		//TODO-DONE: move reader from here

		dataList := service.GenerateData(opSize.Insert, personnelInfo)

		populator := createPopulator(numRecords, streamInsert)
		populator.PopulateData(ctx, client, dataList, opSize)
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
		populator = domain.NewStreamInsert(streamInsert)
	} else {
		populator = domain.NewBulkInsert(bulkInsert)
	}
	return populator
}

func getNoOfOperations(streamInsert int, numRecords int) int {
	if streamInsert > 0 {
		return streamInsert
	}
	return numRecords
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
