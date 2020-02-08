package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", write)
    http.ListenAndServe(":8080", nil)
}

func write(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"HEll", r.URL.Path[1:])
}
