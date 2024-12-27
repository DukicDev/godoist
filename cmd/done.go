package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done <task-id>",
	Short: "Mark a task as completed",
	Long: `Close or mark a task as completed using its task ID.

You can find the task ID by using the "list" command. This command will remove the task from the active list of tasks.`,
	Args: cobra.ExactArgs(1), // Enforces exactly one argument
	Example: `  godoist done 123
  godoist done 42`,
	Run: doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneRun(cmd *cobra.Command, args []string) {
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid task ID: %v. Task ID must be a number.", err)
	}
	taskContent, err := client.CloseTask(index, cacheFile)
	if err != nil {
		log.Fatalf("error while closing task: %v\n", err)
	}
	fmt.Printf("Successfully closed task: %s\n", taskContent)
}
