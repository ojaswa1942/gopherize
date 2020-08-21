package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/ojaswa1942/gopherize/task/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as done",
	Run: do,
}

func do(cmd *cobra.Command, args []string) {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse argument:", arg)
		} else {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		exit(fmt.Sprint("No valid arguments to proceed"), true)
	}

	tasks, err := db.ReadTasks()
	if err != nil {
		exit(fmt.Sprint("Something went wrong:", err), true)
	}

	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		err := db.DeleteTask(task.Key)
		if err != nil {
			fmt.Printf("Failed to mark %d as complete. Error: %s\n", id, err)
		} else {
			fmt.Printf("WooHoo! Marked task %d as complete.\n", id)
		}
	}
}

func init() {
	rootCmd.AddCommand(doCmd)
}
