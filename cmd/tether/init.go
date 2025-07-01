package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new tether configuration file.",
	Long:  `Creates a config.yaml file in the current directory with default values.`,
	RunE:  runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter local path: ")
	localPath, _ := reader.ReadString('\n')
	localPath = strings.TrimSpace(localPath)

	fmt.Print("Enter remote path: ")
	remotePath, _ := reader.ReadString('\n')
	remotePath = strings.TrimSpace(remotePath)

	fmt.Print("Enter remote host (user@server): ")
	remoteHost, _ := reader.ReadString('\n')
	remoteHost = strings.TrimSpace(remoteHost)

	config := struct {
		LocalPath  string `yaml:"localPath"`
		RemotePath string `yaml:"remotePath"`
		RemoteHost string `yaml:"remoteHost"`
	}{
		LocalPath:  localPath,
		RemotePath: remotePath,
		RemoteHost: remoteHost,
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to yaml: %w", err)
	}

	if err := os.WriteFile("config.yaml", data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Println("config.yaml created successfully.")
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}