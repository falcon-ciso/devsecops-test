package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    data, _ := ioutil.ReadFile(r.URL.Query().Get("file"))
    fmt.Fprint(w, string(data))
}
