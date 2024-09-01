package main

import (
    "fmt"
    "net/http"
    "time"
    "log"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from PSPDFKit Engineer!")
    })

    server := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Fatal(server.ListenAndServe())
}