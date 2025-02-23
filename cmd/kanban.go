package cmd

import (
    "fmt"
    . "go-todo-cli/models"
    . "go-todo-cli/utils"
    "github.com/spf13/cobra"
    "strings"
    "unicode/utf8"
)

var width = 40

// kanbanCmd représente la commande "kanban"
var kanbanCmd = &cobra.Command{
    Use:   "kanban",
    Short: "Affiche les tâches sous forme de tableau Kanban",
    Run: func(cmd *cobra.Command, args []string) {
        print("\033[H\033[2J")
        All = true
        tasks := LoadTasks()
        filtered := FilterTasks(tasks)
        displayKanban(filtered)
    },
}

func groupTasksByStatus(tasks []Task) map[TaskStatus][]Task {
    grouped := make(map[TaskStatus][]Task)
    for r, task := range tasks {
        if r < 10 {
            grouped[task.Status] = append(grouped[task.Status], task)
        }
    }

    return grouped
}

func displayKanban(tasks []Task) {
    groupedTasks := groupTasksByStatus(tasks)

    // Ordre des colonnes
    statusOrder := []TaskStatus{Stashed, Pending, InProgress, Completed}
    statusTitles := map[TaskStatus]string{
        Stashed:    "📦 \033[1;35mSTASHED\033[0m",
        Pending:    "📝 \033[1;34mTO-DO\033[0m",
        InProgress: "🚧 \033[1;33mIN PROGRESS\033[0m",
        Completed:  "✅ \033[1;32mDONE\033[0m",
    }

    maxRows := 0
    for _, status := range statusOrder {
        if len(groupedTasks[status]) > maxRows {
            maxRows = len(groupedTasks[status])
        }
    }
    fmt.Println("")
    for _, status := range statusOrder {
        fmt.Printf("%-50s", statusTitles[status])
    }
    fmt.Println()

    for range statusOrder {
        fmt.Printf("%-40s", strings.Repeat("─", 35))
    }

    fmt.Println()

    groupedLines := make(map[TaskStatus][]string)
    maxRow := 0
    for key, tasks := range(groupedTasks){
        groupedLines[key] = computeLines(tasks)
        if len(groupedLines[key]) > maxRow {
            maxRow = len(groupedLines[key])
        }
    }

    var finalLines []string

    for i := 0; i < maxRow; i++ {
        currentLine := "";
        for _, status := range statusOrder{
            linesForStatus := groupedLines[status]
            if len(linesForStatus) > i {
                currentLine = currentLine + linesForStatus[i] + "     "
            } else {
                currentLine = currentLine + strings.Repeat(" ", 40) 
            }
        }
        finalLines = append(finalLines, currentLine)
    }

    for _, line := range finalLines{
        fmt.Println(line)
    }

    fmt.Println()
    fmt.Println()
}

func computeLines(tasks []Task) []string {
    var lines []string

    for _,task := range tasks {
        lines = append(lines, strings.Repeat("_", 35))
        var title = fmt.Sprintf("\033[1mTask #%d\033[0m - (%s)", task.Id, task.Category)
        lines = append(lines, "| " + title + strings.Repeat(" ", 40 - utf8.RuneCountInString(title)) + "|")
        lines = append(lines, "|"+ strings.Repeat("~", 33)+"|")
        var result []string
        var tmp = task.Description
        for utf8.RuneCountInString(tmp) > 31 {
            result = append(result, tmp[:31])
            tmp = tmp[31:]
        }
        if len(tmp) > 0 {
            result = append(result, tmp)
        }

        for _, line := range result{
            lineLen := utf8.RuneCountInString(line)
            if len(line) < 31 {
                lines = append(lines, "| " + line + strings.Repeat(" ", 32 - lineLen) + "|")
            }else{
                lines = append(lines, "| " + line + " |")
            }
        }

        lines = append(lines, strings.Repeat("-", 35))
    }

    return lines
}


func init() {
    rootCmd.AddCommand(kanbanCmd)
    kanbanCmd.Flags().StringVarP(&Filter, "filter", "f", "", "returns tasks matching the given filter")
    kanbanCmd.Flags().BoolVarP(&Detail, "detail", "d", false, "provide detailed view for each task")
    kanbanCmd.Flags().BoolVarP(&All, "all", "a", false, "include all status, event completed tasks")
    kanbanCmd.Flags().StringVarP(&Category, "category", "c", "", "returns tasks matching the given category")
    kanbanCmd.Flags().StringVarP(&Status, "status", "s", "", "returns tesks matching the given status (Pending|InProgress|Completed|Stashed)")
}

