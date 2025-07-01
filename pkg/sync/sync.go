package sync

import (
	"fmt"
	"os/exec"

	"github.com/jawn/tether/pkg/config"
)

// Push syncs files from the local machine to the remote machine.
func Push(cfg *config.Config) error {
	fmt.Printf("Syncing from %s to %s:%s\n", cfg.LocalPath, cfg.RemoteHost, cfg.RemotePath)
	cmd := exec.Command("rsync", "-avz", "--delete", cfg.LocalPath, fmt.Sprintf("%s:%s", cfg.RemoteHost, cfg.RemotePath))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync push failed: %s\n%w", string(output), err)
	}
	fmt.Printf("Push successful:\n%s", string(output))
	return nil
}

// Pull syncs files from the remote machine to the local machine.
func Pull(cfg *config.Config) error {
	fmt.Printf("Syncing from %s:%s to %s\n", cfg.RemoteHost, cfg.RemotePath, cfg.LocalPath)
	cmd := exec.Command("rsync", "-avz", "--delete", fmt.Sprintf("%s:%s", cfg.RemoteHost, cfg.RemotePath), cfg.LocalPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync pull failed: %s\n%w", string(output), err)
	}
	fmt.Printf("Pull successful:\n%s", string(output))
	return nil
}

