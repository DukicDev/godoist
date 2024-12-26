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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&duedate, "duedate", "d", "", "set duedate for task")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "set priority for task")
	addCmd.Flags().StringVar(&description, "desc", "", "set description for task")
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("task argument Required. Usage: godoist add 'task name'")
	}
	content := args[0]
	dateString := ""
	if duedate != "" {
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
