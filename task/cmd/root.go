package cmd

import (
	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/ojaswa1942/gopherize/task/db"
	"path/filepath"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "task",
  Short: "Task is a CLI task manager",
  Long: `Task is a utility CLI tool, a task manager, to help you manage your tasks on-the-go. Find more at: https://github.com/ojaswa1942/gopherize/task `,
}

func Execute() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.InitDb(dbPath)
	if err != nil {
		exit(fmt.Sprint("cannot initiate db:", err), true)
	}
	if err := rootCmd.Execute(); err != nil {
		exit(fmt.Sprintln(err), true)
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
