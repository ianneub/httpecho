package main

import (
  "fmt"
  "net/http"
  "time"
)

func longHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Handling %s request for %s from %s..\n", r.Method, r.URL, r.RemoteAddr)
  time.Sleep(5 * time.Minute)
  fmt.Fprintf(w, "Done.")
}
