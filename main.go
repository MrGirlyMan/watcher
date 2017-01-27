package main

import (
	"os"
	"fmt"
	"lookout"
	"sleuth"
	// "looker"
	// "spotter"
)

var Name = "lookout"

func main() {
  fmt.Println(TEMPLATE_DIR)

	if l := lookout.NewLookout("Sunglasses", "ls"); l == nil {
		os.Exit(1)
	}

  // Launch Sleuth webserver
  sleuth.Sleuth()

	fmt.Println(os.Geteuid())
	fmt.Println(Name)
}
