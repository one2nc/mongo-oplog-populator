package cmd

import (
	"fmt"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/populator"
	"os"

	"github.com/spf13/cobra"
)

var bulkInsert int
var streamInsert int

var rootCmd = &cobra.Command{
	Use:   "mongopop",
	Short: "A simple CLI application",
	Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().HasFlags() {
			cmd.Usage()
			return
		}
		//get Client for mongo
		client := config.GetClient()
		if streamInsert > 0 {
			populator.StreamInsert(client, streamInsert)
		} else if bulkInsert > 0 {
			result := populator.Populate(client, bulkInsert)
			for i := 0; i < len(result); i++ {
				fmt.Printf("result: %v\n", result[i])
			}
		}
		//Disconnect Client
		config.DisconnectClient(client)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&bulkInsert, "bulk", "b", 0, "Bulk Insert")
	rootCmd.Flags().IntVarP(&streamInsert, "stream", "s", 0, "Stream Insert")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
