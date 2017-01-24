package lookout

import (
	"fmt"
)

type Lookout struct {
	Name	string
	Process	string
}

func init() {
	fmt.Println("init() is always called when lookout is imported")
}

func NewLookout(n string, p string) *Lookout {
	lookout := Lookout{Name: n, Process: p}
	fmt.Println(lookout)

	return &lookout
}