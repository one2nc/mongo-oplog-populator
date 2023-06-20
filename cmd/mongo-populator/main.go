package main

import (
	"context"
	"fmt"
	"mongo-oplog-populator/config"
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
	modeFlag bool
)

func init() {
	rootCmd.Flags().BoolVarP(&modeFlag, "stream", "s", false, "Stream Mode")

}

var rootCmd = &cobra.Command{
	Use:   "mongopop",
	Short: "Data Populator in MongoDB",
	Long:  "This utility facilitates fake random data population in a MongoDB by providing functionalities to perform insert, update, and delete operations.",
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().HasFlags() {
			cmd.Usage()
			return
		}

		noOfOperations, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid argument. Please provide a valid integer.")
			os.Exit(1)
		}

		cfg := config.Load()
		var fakeData generator.FakeData
		//checking if csv file already exists, if not, then create a writer
		//TODO: pass cancel context to writer and reader, so that we can terminate
		_, err = os.Stat(cfg.CsvFileName)
		if os.IsNotExist(err) {
			gofakeitGenerator := generator.NewGoFakeItGenerator()
			fakeData = gofakeitGenerator.GenerateFakeData()

			csvWriter := writer.NewCSVWriter(cfg.CsvFileName)
			csvWriter.WriteData(fakeData)
		} else {
			csvReader := reader.NewCSVReader(cfg.CsvFileName)
			fakeData = csvReader.ReadData()
		}

		ctx, cancel := context.WithCancel(context.Background())
		// Handle interrupt signal
		handleInterruptSignal(cancel)

		//get Client for mongo
		client := writer.NewMongoClient(ctx, cfg)

		//Disconnect Client (pass ctx later for disconnecting mongo client)
		defer writer.DisconnectMongoClient(ctx, client)

		populatorService := service.NewPopulator(client, modeFlag, noOfOperations)
		populatorService.PopulateData(ctx, fakeData)
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

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
