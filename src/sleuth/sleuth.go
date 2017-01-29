package sleuth

import (
    "os"
    "fmt"
)

// Slueth is the backend server answering REST calls from frontman.
// Go natively supports many HTTP functions, so researching the native go packages would be a great beginning.

func Sleuth() {
    err := 1
    
    StartHandlers()

    defer fmt.Println("webserver down")

    if (err != 1) {
        os.Exit(1)
    }
}
