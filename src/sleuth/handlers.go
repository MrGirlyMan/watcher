package sleuth

import (
  "os"
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) (*Page, error) {
  var index string

  home_directory := os.Getenv("GOPATH")

  if (home_directory != "") {
    index = home_directory + "/src/frontman/index.html"
  } else {
    log.Fatal("$GOPATH is not define.")
  }

  body, err := ioutil.ReadFile(index)

  if err != nil {
    return nil, err
  }

  return &Page{Title: title, Body: body}, nil
}

// Handlers handle web responses
func handler(w http.ResponseWriter, r *http.Request) {
  p, _ := loadPage("Title")

  fmt.Fprintf(w, string(p.Body))
  // log.Fatal(err)

  // if (err != nil) {
  //   fmt.Fprintf(w, string(p.Body))
  // } else {
  //   fmt.Fprintf(w, "Error reading index.html")
  // }

}

func StartHandlers() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", http.FileServer(http.Dir("./src/frontman/public")))
}
