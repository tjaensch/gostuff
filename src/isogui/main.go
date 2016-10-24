package main

import (
  "log"
  "net/http"
)

func main() {
  router := NewRouter()

  log.Fatal(http.ListenAndServe("10.90.235.15:1313", router))
}
