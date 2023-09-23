package cmd

import (
	"fmt"
	"gabrieleromanato/task/tasks"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `Add a new task to your TODO list`,
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		if description == "" {
			fmt.Println("Please provide a description for the task")
			os.Exit(1)
		}
		tasksData, err := tasks.LoadFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newTask := tasks.Task{
			ID:          tasksData.GetNextID(),
			Description: description,
			Completed:   false,
		}
		tasksData.AddTask(newTask)
		err = tasks.SaveToFile(tasksData)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
