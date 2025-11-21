package main

import (
  "fmt"
  "net/http"
  "time"
  "strconv"
)

func longHandler(w http.ResponseWriter, r *http.Request) {
  qs := r.URL.Query().Get("s")
  seconds, err := strconv.Atoi(qs)
  if (err != nil) {
    seconds = 300
  }

  fmt.Printf("Handling %s request for %s from %s..\n", r.Method, r.URL, r.RemoteAddr)
  time.Sleep(time.Duration(seconds) * time.Second)
  fmt.Fprintf(w, "Done.")
}
