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
    if err != nil {
        fmt.Fprintf(w, "There was an error from the API request %s", err.Error())
    } else {
        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Fprintf(w, "There was an error from parsing the request body %s", err.Error())
        } else {
            fmt.Fprintf(w, string(responseData))
        }
    }
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
