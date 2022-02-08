package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
    log.Println("anonymous func")
    w.WriteHeader(http.StatusOK)
  })
  http.HandleFunc("/", myHandler)
  http.HandleFunc("/m", myMiddleware(myHandler))

  http.ListenAndServe(":4848", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("my handler")
  w.WriteHeader(http.StatusOK)
}

func myMiddleware(h http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("middleWare func")
    h(w, r)
  }
}
