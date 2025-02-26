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
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        masterConfig := LoadMaster()
        category := "None"
        if masterConfig.DefaultCategory != "" {
            category = masterConfig.DefaultCategory
        }
        nextId := masterConfig.LastId + 1
        content, err := OpenTaskEditor("Title", category, "")

        if err != nil {
            fmt.Println("Error while editing task")
            return
        }

        editedTask := ParseTask(content)

        if editedTask.Title == "" {
            fmt.Printf("\n Error : Title must not be empty")
        }

        var completed time.Time
        status := Pending

        if trace {
            completed = time.Now()
            status = Completed
        }

        task := Task{
            Id:          nextId,
            Title:       editedTask.Title,
            Description: editedTask.Description,
            CreatedAt:   time.Now(),
            CompletedAt: completed,
            Category:    editedTask.Category,
            Status:      status,
        }

        tasks := LoadTasks()
        tasks = append(tasks, task)
        SaveTasks(tasks)
        masterConfig.LastId = nextId
        SaveMaster(masterConfig)
        fmt.Println()
        fmt.Println("✅ New task added : \"", task.Id, " - ", task.Title, "\"")
        fmt.Println()
    },

}
var trace bool


func init() {
    addCmd.Flags().BoolVarP(&trace, "trace", "t", false, "Add a task but just to trace it exists, by setting it to Completed immediatly")
    rootCmd.AddCommand(addCmd)
}
