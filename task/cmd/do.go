package cmd

import (
	"fmt"
	"gabrieleromanato/task/tasks"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  `Mark a task on your TODO list as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		if taskID == "" {
			fmt.Println("Please provide a task ID")
			os.Exit(1)
		}
		id, err := strconv.Atoi(taskID)
		if err != nil {
			fmt.Println("Please provide a valid task ID")
			os.Exit(1)
		}
		tasksData, err := tasks.LoadFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		task := tasksData.GetTask(id)
		if task.ID == 0 {
			fmt.Println("Task not found")
			os.Exit(1)
		}
		tasksData.CompleteTask(task)
		err = tasks.SaveToFile(tasksData)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Marked \"%s\" as completed.\n", task.Description)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
