/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gabrieleromanato/task/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasksData, err := tasks.LoadFile()
		if err != nil {
			fmt.Println(err)
		}
		totalTasks := len(tasksData.Tasks)
		if totalTasks == 0 {
			fmt.Println("You have no tasks to complete!")
		} else {
			uncompleted := tasksData.GetUncompletedTasks()
			totalUncompleted := len(uncompleted)
			if totalUncompleted == 0 {
				fmt.Println("You have no tasks to complete!")
			} else {
				fmt.Printf("You have %d tasks to complete:\n", totalUncompleted)
				for _, task := range uncompleted {
					fmt.Printf("%d. %s\n", task.ID, task.Description)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
