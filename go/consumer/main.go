package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    response, err := http.Get(os.Getenv("URL"))
    responseData, err := ioutil.ReadAll(response.Body)
    fmt.Println(string(responseData))
}
