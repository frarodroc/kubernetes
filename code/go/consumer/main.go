package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    url := os.Getenv("URL")
    response, err := http.Get(url)
    w.Header().Set("My-Custom-Header", "1234")
    fmt.Fprintf(w, string(responseData))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
