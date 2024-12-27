package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var duedate string
var priority int
var description string
var dueToday bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new todo",
	Long: `Add a new todo task with optional settings for duedate, priority, and description.

You can set the duedate with the --due (-d) flag or mark the task as due today with --today (-t).
The priority can be set with --priority (-p), and additional details can be added with --desc.`,
	Args:    cobra.ExactArgs(1),
	Example: "godoist add 'Buy groceries' -d 31.12.2024 -p 2 --desc 'Shopping list'",
	Run:     addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&duedate, "due", "d", "", "Set duedate for task")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "Set priority for task")
	addCmd.Flags().StringVar(&description, "desc", "", "Set description for task")
	addCmd.Flags().BoolVarP(&dueToday, "today", "t", false, "Set duedate for task as today")
}

func addRun(cmd *cobra.Command, args []string) {
	content := args[0]
	dateString := ""
	if dueToday {
		dateString = time.Now().Format("2006-01-02")
	} else if duedate != "" {
		date, err := time.Parse("2.1.2006", duedate)
		if err != nil {
			log.Fatalf("could not parse duedate: %v\n", err)
		}
		dateString = date.Format("2006-01-02")
	}

	err := client.CreateTask(content, dateString, priority, description)
	if err != nil {
		log.Fatalf("task creation failed: %v\n", err)
	}
	fmt.Println("Successfully created task!")
}
