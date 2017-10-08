package main

import "net/http"

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/-/", fileServerHandler)
  http.ListenAndServe(":8080", nil)
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
  h := http.FileServer(noDir{http.Dir("public")})
  http.StripPrefix("/-", h).ServeHTTP(w, r)
}

type noDir struct {
  http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
  f, err := Dir.Open(name)
  if err != nil {
    return nil, err
  }
  stat, err := f.Stat()
  if err != nil {
    return nil, err
  }
  if stat.IsDir() {
    return nil, os.ErrNotExist()
  }
  return f, nil
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
