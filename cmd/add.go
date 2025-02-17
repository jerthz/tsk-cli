package cmd

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
    "github.com/spf13/cobra"
)

import . "go-todo-cli/models"
import . "go-todo-cli/utils"

// addCmd represents the add command
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
        saveTasks(tasks)
        masterConfig.LastId = nextId
        SaveMaster(masterConfig)
        fmt.Println("✅ New task added : \"", task.Description, "\"")

    },
}

func saveTasks(tasks []Task) {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    file, err := os.OpenFile(GetTaskFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier :", err)
        return
    }
    defer file.Close()

    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
    } else {
        fmt.Println("Tâche sauvegardée avec succès !")
    }
}


func init() {
    addCmd.Flags().StringVarP(&category, "category", "c", "", "Add this task inside the given category")
    rootCmd.AddCommand(addCmd)
}
