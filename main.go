package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    fmt.Println("Server started at http://localhost:80")
    http.ListenAndServe(":80", nil)
}
