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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	Run:   deleteRun,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("Argument required. Usage: godoist delete <task-id>")
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Argument Error: %v\n", err)
	}
	resp, err := client.DeleteTask(index, cacheFile)
	if err != nil {
		log.Fatalf("Error while trying to delete task: %v\n", err)
	}
	fmt.Print(resp)
}
