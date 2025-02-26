package cmd

import (
	"github.com/spf13/cobra"
	"time"
)
import . "go-todo-cli/utils"
import . "go-todo-cli/models"

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Display tasks started or ended yesterday for your 'daily'",
	Long:  `Display tasks started or ended yesterday for your 'daily'`,
	Run: func(cmd *cobra.Command, args []string) {
		MasterInit()
		print("\033[H\033[2J")
		All = true
		tasks := LoadTasks()
		filtered := FilterTasks(tasks)
		filteredDaily := filterDaily(filtered)
		grouped := GroupByCategory(filteredDaily)
        DisplayInline(grouped)
	},
}

func filterDaily(tasks []Task) []Task {
	var filteredDaily []Task
	startYesterday := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)
	endYesterday := startYesterday.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	for _, task := range tasks {
		if (!task.StartedAt.IsZero() &&
			(task.StartedAt.After(startYesterday) && task.StartedAt.Before(endYesterday))) ||
			(!task.CompletedAt.IsZero() &&
				(task.CompletedAt.After(startYesterday) && task.CompletedAt.Before(endYesterday))) {
			filteredDaily = append(filteredDaily, task)
		}
	}
	return filteredDaily
}

func init() {
	dailyCmd.Flags().StringVarP(&Filter, "filter", "f", "", "returns tasks matching the given filter")
	dailyCmd.Flags().BoolVarP(&Detail, "detail", "d", false, "provide detailed view for each task")
	dailyCmd.Flags().BoolVarP(&All, "all", "a", false, "include all status, event completed tasks")
	dailyCmd.Flags().StringVarP(&Category, "category", "c", "", "returns tasks matching the given category")
	dailyCmd.Flags().StringVarP(&Status, "status", "s", "", "returns tesks matching the given status (Pending|InProgress|Completed|Stashed)")
	rootCmd.AddCommand(dailyCmd)
}
