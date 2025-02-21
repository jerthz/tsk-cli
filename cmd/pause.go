
package cmd

import (
	"fmt"
    "strconv"
	"github.com/spf13/cobra"
    . "go-todo-cli/models"
    . "go-todo-cli/utils"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a task",
	Long: `Pause a task by setting it to Stashed status.
    `,
	Run: func(cmd *cobra.Command, args []string) {
	    MasterInit()

        id, err := strconv.Atoi(args[0])

        if err != nil{
            fmt.Printf("\nError : The task id is not an integer\n\n")
            return
        }

        tasks := LoadTasks()

        for i := range tasks{
            if tasks[i].Id == id {
                if tasks[i].Status == Completed {
                   fmt.Printf("\nThe task is already completed\n\n") 
                   return
                }
                tasks[i].Status = Stashed
                SaveTasks(tasks)
                fmt.Printf("\nTask successfully paused ! \n\n")
                return
            }
        }
        fmt.Printf("\n Unable to find the task with the given id...  \n\n")
    },
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
