package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "task",
  Short: "Task is a CLI task manager",
  Long: `Task is a utility CLI tool, a task manager, to help you manage your tasks on-the-go. Find more at: https://github.com/ojaswa1942/gopherize/task `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
