/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Close a Task",
	Run:   doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doneRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("Argument required. Usage: godoist done <task-id>")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("error while processing arguments: %v\n", err)
	}
	taskContent, err := client.CloseTask(index, cacheFile)
	if err != nil {
		log.Fatalf("error while closing task: %v\n", err)
	}
	fmt.Printf("Successfully closed task: %s\n", taskContent)
}
