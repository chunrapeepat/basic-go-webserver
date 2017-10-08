package main

import "net/http"

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/-/", fileServerHandler) 
  http.ListenAndServe(":8080", nil)
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
  h := http.FileServer(http.Dir("public"))
  http.StripPrefix("/-", h).ServeHTTP(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8 ")
  w.Write([]byte(`
    <!doctype html>
    <title>Chun Rapeepat</title>
    <link rel=stylesheet href=/-/css/design.css />
    <p class="red">
      Static html with GO
    </p>
  `))
}
