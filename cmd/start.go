/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/internal/storage"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		id, err := strconv.Atoi(args[0])
		if err != nil || id <= 0{
			fmt.Println("PLease provide valid task id")
			return
		}
		tasks, err := storage.LoadTask()
		if err != nil{
			fmt.Println("Error loading tasks:", err)
			return
		}	
		for i, task := range tasks{
			if task.ID == id{
				tasks[i].IsRunning = true
				tasks[i].StartTime = time.Now()
				err = storage.SaveTask(tasks)
				if err != nil {
					fmt.Println("Error saving tasks:", err)
					return
				}
				fmt.Printf("Task %d started\n", id)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
