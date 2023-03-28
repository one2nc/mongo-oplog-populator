package cmd

import (
    "fmt"
    "mongo-oplog-populator/config"
    "mongo-oplog-populator/populator"
    "os"

    "github.com/spf13/cobra"
)

var operations int
var batchInsert int

var rootCmd = &cobra.Command{
    Use:   "oplogpop",
    Short: "A simple CLI application",
    Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
    Run: func(cmd *cobra.Command, args []string) {
        //get Client for mongo
        client := config.GetClient()
        if batchInsert > 0 {
            populator.BatchInsert(client, batchInsert)
        } else {
            result := populator.Populate(client, operations)
            for i := 0; i < len(result); i++ {
                fmt.Printf("result: %v\n", result[i])
            }
        }
        //Disconnect Client
        config.DisconnectClient(client)
    },
}

func init() {
    rootCmd.Flags().IntVar(&operations, "op", 1, "No of operations to perform")
    rootCmd.Flags().IntVar(&batchInsert, "b", 0, "Batch Insert")
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}