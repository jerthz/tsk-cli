package cmd

import (
	"github.com/spf13/cobra"
)

// Définition de la commande trace comme alias de add -t
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Shortcut for adding a trace task",
	Long:  `Shortcut for adding a task with the trace flag (-t).`,
	Run: func(cmd *cobra.Command, args []string) {
		trace = true
		addCmd.Run(cmd, args)
	},
}



func init() {
    rootCmd.AddCommand(traceCmd)
}