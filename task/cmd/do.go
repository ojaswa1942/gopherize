package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
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

	fmt.Println("do called", ids)
}

func init() {
	rootCmd.AddCommand(doCmd)
}
