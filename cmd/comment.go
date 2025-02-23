package cmd
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
    "strconv"
)

import . "go-todo-cli/utils" 

var commentCmd = &cobra.Command{
	Use:   "comment [task ID]",
	Short: "Modifie la description d'une tâche",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        
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

		tmpFile, err := os.CreateTemp("", "task_comment_*.txt")
		if err != nil {
            fmt.Println("\nError : Unable to create temp file to edit the task")
			return
		}
		defer os.Remove(tmpFile.Name()) 

		tmpFile.WriteString(task.Comment)
		tmpFile.Close()

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nvim" // Par défaut sur la plupart des Unix
		}
		cmdEdit := exec.Command(editor, tmpFile.Name())
		cmdEdit.Stdin = os.Stdin
		cmdEdit.Stdout = os.Stdout
		cmdEdit.Stderr = os.Stderr

		err = cmdEdit.Run()
		if err != nil {
			fmt.Println("❌ Error while editing the task")
			return
		}

		content, err := ioutil.ReadFile(tmpFile.Name())
		if err != nil {
			fmt.Println("❌ Impossible to read temp file after editing")
			return
		}

		task.Comment = strings.TrimSpace(string(content))
		err = UpdateTask(task)
		if err != nil {
			fmt.Println("❌ Impossible to update the task...")
			return
		}

		fmt.Printf("\n✅ Comment successfully updated !\n\n")
	},
}

func init() {
	rootCmd.AddCommand(commentCmd)
}

