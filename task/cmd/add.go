package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
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
	fmt.Println("called add", task)
}

func init() {
	rootCmd.AddCommand(addCmd)
}