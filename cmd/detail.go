/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"
    "strconv"
    "github.com/spf13/cobra"
)
import . "go-todo-cli/utils"
import . "go-todo-cli/models"

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
    Use:   "detail",
    Short: "Show everything about a task, including timestamps and comments",
    Long: `A complete view of a task, including status, start date, end date, comments..
    `,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        print("\033[H\033[2J")
        id, err := strconv.Atoi(args[0])

        if err != nil {
            fmt.Printf("\nError : The task id is not an integer !\n\n")
            return
        }

        tasks := LoadTasks()
        for i := range tasks {
            if tasks[i].Id == id {
                displayDetailedTask(tasks[i])
                return
            }
        }

        fmt.Printf("\n Unable to find the task with the given id ...")
    },
}

func displayDetailedTask(task Task) {
    fmt.Printf("\n\033[1mTask Id :\033[0m %d \n", task.Id) 
    fmt.Printf("\033[1mTask Title :\033[0m %s \n", task.Title) 
    fmt.Printf("\033[1mTask Status :\033[0m %s \n", task.Status.String())

    if task.Description != "" {
        fmt.Printf("\n\033[1mComment :\033[0m\n\n%s \n", task.Description)
    }
    fmt.Println("\n")

}

func init() {
    rootCmd.AddCommand(detailCmd)
}
