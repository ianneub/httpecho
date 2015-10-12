package main

import (
  "fmt"
  "net/http"
  "strings"
  "sort"
)

func sortMapByKeys(input map[string][]string) (keys []string) {
  // To store the keys in slice in sorted order
  for k := range input {
    keys = append(keys, k)
  }
  sort.Strings(keys)

  return
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Handling %s request for %s from %s..\n", r.Method, r.URL, r.RemoteAddr)

  for _, name := range sortMapByKeys(r.Header) {
    fmt.Fprintf(w, "%s: %s\n", name, strings.Join(r.Header[name], ""))
  }
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Listening on port 8080..")
  http.ListenAndServe(":8080", nil)
}
