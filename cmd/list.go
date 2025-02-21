package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "strings"
)

import . "go-todo-cli/models"
import . "go-todo-cli/utils"

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Long: `List all tasks. Only returns pending ones by default
    `,
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        tasks := LoadTasks()
        filteredTasks := filterTasks(tasks)
        if len(filteredTasks) == 0 {
            fmt.Printf("\n📭 No task registered.\n\n")
            return
        }
        grouped := groupByCategory(filteredTasks)

        for category, taskList := range grouped {
            fmt.Printf("\n\n")
            c := category
            if c == "" {
                c = "No category"
            }
            fmt.Printf("\033[1m %s :\033[0m\n\n", c)
            for _, task := range taskList {
                fmt.Printf("    %d. %s (%s)\n", task.Id, task.Description, task.Status.String())
            }
        }
        fmt.Printf("\n\n")
    },
}

func filterTasks(tasks []Task) []Task{
    var filteredCategory []Task
    for _, task := range tasks {
        if category == "" || strings.Contains(task.Category,category) {
            filteredCategory = append(filteredCategory, task)
        }
    }
    var filteredText []Task
    for _, task := range filteredCategory{
        if filter == "" || strings.Contains(task.Description,filter) {
            filteredText = append(filteredText, task)
        }
    }  
    var filteredStatus []Task
    for _, task := range filteredText {
        if all || strings.ToLower(status) == "completed" || task.Status != Completed {
            filteredStatus = append(filteredStatus, task)
        }
    }

    var filteredStatus2 []Task
    for _, task := range filteredStatus {
        if status != "" {
            taskStatus, err := StringToTaskStatus(status)
            if err != nil{
                fmt.Printf("\nThe given status does not exist, please use Pending|InProgress|Completed|Stashed\n\n")
            }

            if taskStatus == task.Status {
                filteredStatus2 = append(filteredStatus2, task)
            }
        } else{
            filteredStatus2 = append(filteredStatus2, task)
        }
    }

    return filteredStatus2
}

func groupByCategory(tasks []Task) map[string][]Task {
    grouped := make(map[string][]Task)
    for _, task := range tasks {
        grouped[task.Category] = append(grouped[task.Category], task)
    }

    return grouped
}

var filter string
var category string
var detail bool
var all bool
var status string

func init() {
    listCmd.Flags().StringVarP(&filter, "filter", "f", "", "returns tasks matching the given filter")
    listCmd.Flags().BoolVarP(&detail, "detail", "d", false, "provide detailed view for each task")
    listCmd.Flags().BoolVarP(&all, "all", "a", false, "include all status, event completed tasks")
    listCmd.Flags().StringVarP(&category, "category", "c", "", "returns tasks matching the given category")
    listCmd.Flags().StringVarP(&status, "status", "s", "", "returns tesks matching the given status (Pending|InProgress|Completed|Stashed)")
    rootCmd.AddCommand(listCmd)
}
