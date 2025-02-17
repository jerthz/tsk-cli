/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

import . "go-todo-cli/utils"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a task",
	Long: `Start a task by setting it to Pending status. Also register a start datetime
`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        tasks := LoadTasks()
             
    },
}

func init() {
	rootCmd.AddCommand(startCmd)
}
