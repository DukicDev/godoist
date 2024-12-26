/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var taskMaxLength int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Todos",
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().IntVarP(&taskMaxLength, "length", "l", 50, "set max length for task string")
}

func listRun(cmd *cobra.Command, args []string) {
	tasks, err := client.GetTasks()
	if err != nil {
		log.Fatalf("error while listing Todos: %v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintln(w, "ID\tTask\tDescription\tDue Date")

	for i, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.ShortContent(taskMaxLength), task.Description, task.DueDate)
	}

	w.Flush()
}