package main

import (
    "os"
    // "fmt"

    "watcher/config"
    "watcher/lookout"
    "watcher/sleuth"
)

var Name = "lookout"

func main() {
    // For now, assume settings.yml is in the same directory as main.go
    // TODO: Add command line arguments for settings files
    configs := config.Configs("settings.yml")

    if configs == nil {
        os.Exit(1)
    }

    // Launch Lookout process and network monitoring
    if err := lookout.Run(configs); err != nil {
        os.Exit(1)
    }

    // Launch Sleuth webserver
    sleuth.Sleuth()
}
