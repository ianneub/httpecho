package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/long", longHandler)
  http.HandleFunc("/", echoHandler)
  fmt.Println("Listening on port 8080..")
  http.ListenAndServe(":8080", nil)
}
