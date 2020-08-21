package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"github.com/ojaswa1942/gopherize/task/db"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a task to task list",
  Run: add,
}

func add (cmd *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	if task == "" {
		exit(fmt.Sprint("Argument for command expected, cannot add"), true)
	}

	_, err := db.CreateTask(task)
	if err != nil {
		exit(fmt.Sprint("Looks like fortune wants you to rest! Some error occurred:", err), true)
	}
	fmt.Printf("Added task `%s`\n", task)
}

func init() {
	rootCmd.AddCommand(addCmd)
}