package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("client/static"))
  http.Handle("/", http.StripPrefix("client", fs))

  port := ":3000"
  log.Printf("Listening on port %s\n", port)
  http.ListenAndServe(port, nil)
}