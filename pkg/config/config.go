package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the configuration file.
type Config struct {
	LocalPath  string `yaml:"localPath"`
	RemotePath string `yaml:"remotePath"`
	RemoteHost string `yaml:"remoteHost"`
}

// LoadConfig reads the configuration from the given path.
func LoadConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file not found at %s. Please run 'tether init'", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling configuration: %w", err)
	}

	return &cfg, nil
}
