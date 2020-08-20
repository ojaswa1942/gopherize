package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of all pending tasks",
	Run: list,
}

func list (cmd *cobra.Command, args []string) {
	fmt.Println("list called")
}

func init() {
	rootCmd.AddCommand(listCmd)
}
