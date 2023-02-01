package main

import (
    "fmt"
    "log"
    "net/http"
)

func main(){
    http.HandleFunc("/", homePage)
    fmt.Fprintf(w, "Hello World!")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
