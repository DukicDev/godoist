package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var taskMaxLength int
var showProjects bool
var filter string
var all bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Todos",
	Long: `List all active tasks from your Todoist account.

You can use filters to narrow down tasks (e.g., today, overdue) and show additional details like project names.
Tasks can also be shortened to a specific length for better visibility in the terminal.`,
	Args: cobra.NoArgs,
	Example: `  godoist list --filter today
  godoist list --all --show-projects
  godoist list --length 30`,
	Run: listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&taskMaxLength, "length", "l", 50, "Set max length for task string")
	listCmd.Flags().BoolVar(&showProjects, "show-projects", false, "Show project names for tasks")
	listCmd.Flags().StringVarP(&filter, "filter", "f", "(today|overdue)", "Filter tasks using Todoist filters")
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all active tasks (ignores filters)")
}

func listRun(cmd *cobra.Command, args []string) {
	if all {
		filter = ""
	}
	tasks, err := client.GetTasks(cacheFile, showProjects, filter)
	if err != nil {
		log.Fatalf("error while listing Todos: %v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	if showProjects {
		fmt.Fprintln(w, "ID\tPriority\tTask\tDescription\tDue Date\tProject")

		for i, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n", i+1, task.GetPriority(), task.ShortContent(taskMaxLength), task.Description, task.GetDate(), task.Project)
		}
	} else {
		fmt.Fprintln(w, "ID\tPriority\tTask\tDescription\tDue Date")

		for i, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", i+1, task.GetPriority(), task.ShortContent(taskMaxLength), task.Description, task.GetDate())
		}
	}

	w.Flush()
}
