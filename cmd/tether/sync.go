package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/jawn/tether/pkg/config"
	"github.com/jawn/tether/pkg/sync"
	"github.com/spf13/cobra"
)

var watch bool

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync files between local and remote",
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push files from local to remote",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		if err := sync.Push(cfg); err != nil {
			log.Fatalf("Push failed: %v", err)
		}

		if watch {
			watchAndPush(cfg)
		}
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull files from remote to local",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		if err := sync.Pull(cfg); err != nil {
			log.Fatalf("Pull failed: %v", err)
		}

		if watch {
			log.Println("Watching for remote changes is not yet implemented.")
		}
	},
}

func watchAndPush(cfg *config.Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					if err := sync.Push(cfg); err != nil {
						log.Printf("Push failed during watch: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(cfg.LocalPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func init() {
	pushCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch for file changes and sync automatically")
	pullCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch for file changes and sync automatically")
	syncCmd.AddCommand(pushCmd)
	syncCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(syncCmd)
}