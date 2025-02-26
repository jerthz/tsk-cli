package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

import . "go-todo-cli/utils"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change settings of the cli",
	Long: `Change settings like the default category used
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		MasterInit()
		parts := strings.SplitN(args[0], "=", 2)
		if parts[0] == "defaultCategory" {
			masterConfig := LoadMaster()
			masterConfig.DefaultCategory = parts[1]
			SaveMaster(masterConfig)
			fmt.Println()
			fmt.Println("✅  Default Cagegory updated to", masterConfig.DefaultCategory)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
