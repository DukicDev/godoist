package cmd

import (
	"errors"
	"fmt"
	"github.com/DukicDev/godoist/todoist"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var client *todoist.Client
var cacheFile string

var rootCmd = &cobra.Command{
	Use:   "godoist",
	Short: "A CLI for Todoist",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token, err := GetApiToken()
		if err != nil {
			log.Fatalln("Please set the TODOIST_API_TOKEN environment variable")
		}
		client = todoist.NewClient(token)
	},
	Run: listRun,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("User homedir could not be found. Please use --cache-file")
	}
	rootCmd.PersistentFlags().StringVar(&cacheFile, "cache-file", home+string(os.PathSeparator)+".godoist.json", "set path to cache file")
}

func GetApiToken() (string, error) {
	token := os.Getenv("TODOIST_API_TOKEN")
	if token == "" {
		fmt.Println(`
The environment variable TODOIST_API_TOKEN is not set.

To set it:
- On Linux/Mac: export TODOIST_API_TOKEN=your_todoist_api_token
- On Windows (Command Prompt): set TODOIST_API_TOKEN=your_todoist_api_token
- On Windows (PowerShell): $env:TODOIST_API_TOKEN="your_todoist_api_token"
`)
		return "", errors.New("environment variable TODOIST_API_TOKEN is required")
	}
	return token, nil
}
