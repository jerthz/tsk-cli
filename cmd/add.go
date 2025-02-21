package cmd

import (
    "fmt"
    "time"
    "github.com/spf13/cobra"
)

import . "go-todo-cli/models"
import . "go-todo-cli/utils"

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new Todo task to the list",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        masterConfig := LoadMaster()
        nextId := masterConfig.LastId + 1


        task := Task{
            Id: nextId,
            Description:    args[0],
            CreatedAt:  time.Now(),
            Category:   category,
            Status: Pending,
        }

        tasks := LoadTasks()
        tasks = append(tasks, task)
        SaveTasks(tasks)
        masterConfig.LastId = nextId
        SaveMaster(masterConfig)
        fmt.Println("✅ New task added : \"", task.Description, "\"")

    },
}



func init() {
    addCmd.Flags().StringVarP(&category, "category", "c", "", "Add this task inside the given category")
    rootCmd.AddCommand(addCmd)
}
