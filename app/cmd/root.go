package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codefresh-hello",
	Short: "codefresh-hello is a test for GKE and Codefresh integration",
	Run: func(cmd *cobra.Command, args []string) {
		// Run help by default
		cmd.Help()
	},
}

// Execute runs root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command %s", err)
	}
}
