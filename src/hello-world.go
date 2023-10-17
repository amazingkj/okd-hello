package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    message := os.Getenv("MESSAGE")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
    })

    fmt.Println("Starting server on port 8086. webhook test")
    http.ListenAndServe(":8086", nil)
}
