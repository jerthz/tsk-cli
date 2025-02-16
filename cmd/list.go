/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)


// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Long: `List all tasks. Only returns pending ones by default
    `,
    Run: func(cmd *cobra.Command, args []string) {
        tasks := loadTasks()
        if len(tasks) == 0 {
            fmt.Println("📭 Aucune tâche enregistrée.")
            return
        }

        fmt.Println("\n\033[1m📋 Liste des tâches :\033[0m\n") 
        for i, task := range tasks {
            fmt.Printf("    %d. %s (Ajoutée le %s)\n", i+1, task.Description, task.CreatedAt.Format("02/01/2006 15:04"))
        }
        fmt.Println()
    },
}

var filter string
var category string
var detail bool

func init() {
    listCmd.Flags().StringVarP(&filter, "filter", "f", "", "returns tasks matching the given filter")
    listCmd.Flags().BoolVarP(&detail, "detail", "d", false, "provide detailed view for each task")
    listCmd.Flags().StringVarP(&category, "category", "c", "", "returns tasks matching the given category")
    rootCmd.AddCommand(listCmd)
    
    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // listCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
