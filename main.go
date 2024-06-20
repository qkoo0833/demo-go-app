package main

import (
    "fmt"
    "log"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, after rollback!")
}

func main() {
    http.HandleFunc("/", index)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
