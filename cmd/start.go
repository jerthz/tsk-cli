package cmd

import (
	"fmt"
    "strconv"
    "time"
	"github.com/spf13/cobra"
	. "go-todo-cli/utils"
    . "go-todo-cli/models"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a task",
	Long: `Start a task by setting it to InProgress status. Also register a start datetime
`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()

        id, err := strconv.Atoi(args[0])
        
        if err != nil {
            fmt.Printf("\nError : The task id is not an integer !\n\n")
            return
        }

        tasks := LoadTasks()
        for i := range tasks {
            if tasks[i].Id == id {
                if tasks[i].Status == InProgress{
                    fmt.Printf("\nThe task is already started\n\n")
                    return
                }
                if tasks[i].Status == Completed {
                   fmt.Printf("\nThe task is already completed\n\n") 
                   return
                }
                tasks[i].Status = InProgress
                tasks[i].StartedAt = time.Now()
                SaveTasks(tasks)
                fmt.Printf("\nTask successfully started !\n\n")
                return
            }
        }

        fmt.Printf("\n Unable to find the task with the given id...  \n\n")
    },
}

func init() {
	rootCmd.AddCommand(startCmd)
}
