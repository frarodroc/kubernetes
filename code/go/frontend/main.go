package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/google/uuid"
)

func homePage(w http.ResponseWriter, r *http.Request){
    response, err := http.Get("http://backend.backend:8080")
    id := uuid.New()
    w.Header().Set("My-Custom-Header", id.String())
    fmt.Fprintf(w, string(responseData))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
