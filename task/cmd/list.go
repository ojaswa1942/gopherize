package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ojaswa1942/gopherize/task/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of all pending tasks",
	Run: list,
}

func list (cmd *cobra.Command, args []string) {
	tasks, err := db.ReadTasks()
	if err != nil {
		exit(fmt.Sprint("Cannot read tasks:", err), true)
	}
	if len(tasks) == 0 {
		fmt.Println("No pending tasks! Time to party?")
	} else {
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
