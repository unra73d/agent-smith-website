package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "favicon.ico")
    })

    http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "logo.png")
    })

    fmt.Println("Server started at http://localhost:80")
    http.ListenAndServe(":80", nil)
}
