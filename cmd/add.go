package cmd

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
    "github.com/spf13/cobra"
)

import . "go-todo-cli/models"

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new Todo task to the list",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        task := Task{
            Description:    args[0],
            CreatedAt:  time.Now(),
        }

        tasks := loadTasks()
        tasks = append(tasks, task)
        saveTasks(tasks)

        fmt.Println("✅ New task added : \"", task.Description, "\"")

    },
}

func loadTasks() []Task{
    var tasks []Task
    file, err := os.ReadFile(getTaskFilePath())
    if err == nil {
        json.Unmarshal(file, &tasks)
    }
    return tasks
}

func saveTasks(tasks []Task) {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    file, err := os.OpenFile(getTaskFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func getTaskFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erreur : impossible de trouver le répertoire utilisateur")
		os.Exit(1)
	}
	return home + "/.gotasks.json"
}

func init() {
    rootCmd.AddCommand(addCmd)
}
