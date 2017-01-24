package main

import (
	"os"
	"fmt"
	"lookout"
	// "looker"
	// "spotter"
)

var Name = "lookout"

func main() {
	if l := lookout.NewLookout("Sunglasses", "ls"); l == nil {
		os.Exit(1)
	}

	fmt.Println(os.Geteuid())
	fmt.Println(Name)
}