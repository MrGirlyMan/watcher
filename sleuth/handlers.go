package sleuth

import (
  "os"
  "fmt"
  "log"
  "strings"
  "io/ioutil"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

type Page struct {
  Path string
  Body  []byte
}

func loadPage(filepath string) (*Page, error) {
  var full_path string

  home_directory := os.Getenv("GOPATH")

  if (home_directory != "") {
    full_path = home_directory + filepath
  } else {
    log.Fatal("$GOPATH is not defined.")
  }

  body, err := ioutil.ReadFile(full_path)

  if err != nil {
    return nil, err
  }

  return &Page{Path: full_path, Body: body}, nil
}

// Handlers handle web responses
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  p, _ := loadPage("/src/watcher/frontman/build/index.html")

  fmt.Fprintf(w, string(p.Body))
  // log.Fatal(err)

  // if (err != nil) {
  //   fmt.Fprintf(w, string(p.Body))
  // } else {
  //   fmt.Fprintf(w, "Error reading index.html")
  // }
}


// TODO: Lock this down. Currently someone could access any file on your system by relative paths
func Static(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  var mime_type string

  static_path := "/src/watcher/frontman/build" + params.ByName("filepath")
  p, _ := loadPage(static_path)

  // Determining the MIME type of the file based on file extension
  switch slice := strings.Split(static_path, "."); slice[len(slice) - 1] {
    case "css":
      mime_type = "text/css"
    case "js":
      mime_type = "application/javascript"
    default:
      mime_type = "text/plain"
  }
  
  w.Header().Set("Content-Type", mime_type)

  // Writing content to the http.ResponseWriter
  fmt.Fprintf(w, string(p.Body))
}

func StartHandlers() {
  router := httprouter.New()

  router.GET("/", Index)
  router.GET("/static/*filepath", Static)

  log.Fatal(http.ListenAndServe(":8080", router))
}
