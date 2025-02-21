package cmd

import (
	"fmt"
    "time"
    "strconv"
	"github.com/spf13/cobra"
    . "go-todo-cli/models"
    . "go-todo-cli/utils"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End a task",
	Long: `End a task by setting it to Completed status. Also register a completion DateTime
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
                tasks[i].Status = Completed
                tasks[i].CompletedAt = time.Now()    
                SaveTasks(tasks)
                fmt.Printf("\nTask successfully completed ! \n\n")
                return
            }
        }
        fmt.Printf("\n Unable to find the task with the given id...  \n\n")
    },
}

func init() {
	rootCmd.AddCommand(endCmd)
}
