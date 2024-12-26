package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

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
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatalln("task argument Required. Usage: godoist add 'task name'")
	}
	content := args[0]
	err := client.CreateTask(content)
	if err != nil {
		log.Fatalf("task creation failed: %v\n", err)
	}
	fmt.Println("Successfully created task!")
}
