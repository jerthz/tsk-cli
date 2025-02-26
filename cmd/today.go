package cmd

import (
	"github.com/spf13/cobra"
	"time"
)
import . "go-todo-cli/utils"
import . "go-todo-cli/models"

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Display tasks started or ended today",
	Long:  `Display tasks started or ended today`,
	Run: func(cmd *cobra.Command, args []string) {
		MasterInit()
		print("\033[H\033[2J")
		All = true
		tasks := LoadTasks()
		filtered := FilterTasks(tasks)
		filteredToday := filterToday(filtered)
		grouped := GroupByCategory(filteredToday)
        DisplayInline(grouped)
	},
}

func filterToday(tasks []Task) []Task {
	var filteredDaily []Task
	startToday := time.Now().Truncate(24 * time.Hour)
	endToday := startToday.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	for _, task := range tasks {
		if (!task.StartedAt.IsZero() &&
			(task.StartedAt.After(startToday) && task.StartedAt.Before(endToday))) ||
			(!task.CompletedAt.IsZero() &&
				(task.CompletedAt.After(startToday) && task.CompletedAt.Before(endToday))) {
			filteredDaily = append(filteredDaily, task)
		}
	}
	return filteredDaily
}

func init() {
	todayCmd.Flags().StringVarP(&Filter, "filter", "f", "", "returns tasks matching the given filter")
	todayCmd.Flags().BoolVarP(&Detail, "detail", "d", false, "provide detailed view for each task")
	todayCmd.Flags().BoolVarP(&All, "all", "a", false, "include all status, event completed tasks")
	todayCmd.Flags().StringVarP(&Category, "category", "c", "", "returns tasks matching the given category")
	todayCmd.Flags().StringVarP(&Status, "status", "s", "", "returns tesks matching the given status (Pending|InProgress|Completed|Stashed)")
	rootCmd.AddCommand(todayCmd)
}
