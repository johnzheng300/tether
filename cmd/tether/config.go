package main

import (
	"fmt"

	"github.com/jawn/tether/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var interactive bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or edit the tether configuration.",
	Long:  `View the current configuration or use the -i flag to enter an interactive configuration mode.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if interactive {
			return interactiveConfig()
		}
		return viewConfig()
	},
}

func viewConfig() error {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config to yaml: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func interactiveConfig() error {
	return runInit(nil, nil)
}

func init() {
	configCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Enter interactive configuration mode.")
	rootCmd.AddCommand(configCmd)
}
