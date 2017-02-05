package lookout

import (
	"fmt"
	"log"
	"bytes"

	cfg "watcher/config"
	"gopkg.in/redis.v5"
)

type Lookout struct {
	Name		string
	Process		string
	Cache 		*redis.Client
	config 		*cfg.Config
}

// func init() {
// 	fmt.Println("init() is always called when lookout is imported")
// }


func Run(config *cfg.Config) error {
    return handleError(newLookout("Name", "httpd", config))
}


func newLookout(n string, p string, c *cfg.Config) error {
	l := Lookout{Name: n, Process: p, config: c}
	
	l.attachCache()
	l.inspectProcesses()

	err := l.Cache.Set("redis:key", "5", 0).Err()

	if err != nil {
		panic(err)
	}

	return nil
}

func (l *Lookout) attachCache() {
	l.Cache = redis.NewClient(&redis.Options{
		Addr:		l.config.Redis,
		Password: 	"",
		DB:			0,
	})
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