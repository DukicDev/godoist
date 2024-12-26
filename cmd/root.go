/*
Copyright © 2024 Dino Dukic dukic.dev@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godoist.yaml)")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("User home dir could not be found. Pleae set --cache-file")
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
