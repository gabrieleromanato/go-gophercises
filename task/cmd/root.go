package cmd

import (
	"gabrieleromanato/task/tasks"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI Task Manager",
	Long:  `task is a CLI for managing your TODOs.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if !tasks.FileExists() {
		tasks.CreateFile()
	}
}
