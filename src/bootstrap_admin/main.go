package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func dsmmRatings(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // Form submitted
        r.ParseForm() // Required if you don't call r.FormValue()
        fmt.Println(r.Form["ratings"])
    }
    html_template, _ := ioutil.ReadFile("html_template.html")
    w.Write([]byte(html_template))
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", dsmmRatings)
    http.ListenAndServe("10.90.235.15:1313", nil)
}
