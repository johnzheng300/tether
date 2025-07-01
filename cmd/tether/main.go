package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tether",
	Short: "Tether is a tool for syncing files between a local and remote machine.",
	Long:  `Tether is a CLI tool that provides real-time file synchronization between a local development environment and a remote server.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is provided
		fmt.Println("Welcome to Tether! Use 'tether --help' to see available commands.")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
