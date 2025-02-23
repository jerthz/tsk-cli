package cmd
import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
    "strconv"
)

import . "go-todo-cli/utils" 

var editCmd = &cobra.Command{
	Use:   "edit [task ID]",
	Short: "Edit any task by Id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        print("\033[H\033[2J")

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

        content, err := OpenTaskEditor(task.Title, task.Category, task.Description)

        if err != nil {
            fmt.Println("Error while editing task")
            return
        }

        editedTask := ParseTask(content)


        if editedTask.Title == "" {
            fmt.Printf("\n Error : Title must not be empty")
            return;
        }
           
        task.Title = strings.TrimSpace(editedTask.Title)
        task.Description = strings.TrimSpace(editedTask.Description)
        task.Category = strings.TrimSpace(editedTask.Category)

		err = UpdateTask(task)
		if err != nil {
			fmt.Println("❌ Impossible to update the task...")
			return
		}

		fmt.Printf("\n✅ Comment successfully updated !\n\n")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

