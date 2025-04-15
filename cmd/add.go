/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/internal/storage"
	"cli/internal/task"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var title string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if title == ""{
			fmt.Println("Please provide a title for the task.")
			return
		}
		tasks, err := storage.LoadTask()
		if err !=nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		newTask:= task.Task{
			ID: len(tasks) + 1,
			Title: title,
			IsRunning: false,
			CreatedAt: time.Now(),
		}

		tasks = append(tasks, newTask)
		err = storage.SaveTask(tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task added: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
