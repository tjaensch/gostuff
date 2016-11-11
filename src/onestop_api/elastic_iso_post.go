package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    data, err := os.Open("crap_test_file.xml")
    if err != nil {
        log.Fatal(err)
    }
    defer data.Close()
    req, err := http.NewRequest("POST", "http://localhost:8097/onestop/api/metadata/", data)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/xml")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
}
