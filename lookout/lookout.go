package lookout

import (
	"fmt"
	"log"
	"bytes"

	cfg "watcher/config"
)

type Lookout struct {
	Name	string
	Process	string
	config *cfg.Config
}

// func init() {
// 	fmt.Println("init() is always called when lookout is imported")
// }


func Run(config *cfg.Config) error {
    return handleError(newLookout("Name", "httpd", config))
}


func newLookout(n string, p string, config *cfg.Config) error {
	lookout := Lookout{Name: n, Process: p}

	lookout.inspectProcesses()

	return nil
}



func handleError(err error) error {
	var buf bytes.Buffer

	logger := log.New(&buf, "ErrorLog: ", log.Lshortfile)

	if err == nil {
		return nil
	}

	logger.Print("Exiting: %v", err)
	fmt.Print(&buf)

	return err
}