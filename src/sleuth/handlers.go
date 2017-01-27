package sleuth

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) (*Page, error) {
    filename := title
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// Handlers handle web responses
func handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Home Page: %s", request.URL.Path[1:])
  // page, _ := loadPage("index.html")
  // fmt.Fprintf(writer, page.Body)

}

func networkHandler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Network Page: %s", request.URL.Path[1:])
}

func processHandler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Process Page: %s", request.URL.Path[1:])
}

func StartHandlers() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/network/", networkHandler)
  http.HandleFunc("/process/", processHandler)
  http.ListenAndServe(":8080", nil)
}
