package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    url := os.Getenv("URL")
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("There was an error from the API request %s", err.Error())
    } else {
        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("There was an error from parsing the request body %s", err.Error())
        } else {
            fmt.Println(string(responseData))
        }
    }
}
