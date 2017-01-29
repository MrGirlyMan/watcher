package main

import (
    "os"
    "fmt"
    "lookout"
    "sleuth"
)

var Name = "lookout"

func main() {
    // Launch Lookout process and network monitoring
    if l := lookout.NewLookout("Sunglasses", "ls"); l == nil {
        os.Exit(1)
    }

    // Launch Sleuth webserver
    sleuth.Sleuth()
}
