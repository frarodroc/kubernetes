package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "time"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    response, err := http.Get("http://backend.backend:8080")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    time := time.Now()
    w.Header().Set("My-Custom-Header", time.String())
    fmt.Fprintf(w, string(responseData))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
