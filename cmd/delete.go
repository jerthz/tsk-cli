package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)
import . "go-todo-cli/utils"
import . "go-todo-cli/models"
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		MasterInit()

        id, err := strconv.Atoi(args[0])

        if err != nil {
            fmt.Printf("\nError : The task id is not an integer\n\n")
            return
        }

		task, err := GetTaskById(id)
		if err != nil {
            fmt.Println("\nError : Unable to find the task with the given id... \n\n")
			return
		}

		tasks := removeTaskById(LoadTasks(), task.Id)
		SaveTasks(tasks)
		if err != nil {
			fmt.Println("❌ Impossible to delete the task...")
			return
		}

		fmt.Printf("\n✅ Task successfully deleted !\n\n")

	},
}

func removeTaskById(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.Id == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
