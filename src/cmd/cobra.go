package cmd

import (
	"example.fibodt.com/greetings/src/cmd/greeting"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "greetings",
	Short:             "greetings go",
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(greeting.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(255)
	}
}
