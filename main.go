package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola Mundo desde Go!")
}

func main() {
    http.HandleFunc("/", handler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Puerto por defecto si no se asigna uno
    }
    http.ListenAndServe(":" + port, nil)
}
