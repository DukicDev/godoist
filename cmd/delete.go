package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <task-id>",
	Short: "Delete a task",
	Long: `Delete a task from your Todoist account using its task ID.

The task ID can be obtained by using the "list" command. This command permanently removes the task from your Todoist account.`,
	Args: cobra.ExactArgs(1), // Enforces exactly one argument
	Example: `  godoist delete 123
  godoist delete 42`,
	Run: deleteRun,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteRun(cmd *cobra.Command, args []string) {
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid task ID: %v. Task ID must be a number.", err)
	}
	resp, err := client.DeleteTask(index, cacheFile)
	if err != nil {
		log.Fatalf("Error while trying to delete task: %v\n", err)
	}
	fmt.Print(resp)
}
