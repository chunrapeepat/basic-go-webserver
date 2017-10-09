package main

import (
    "net/http"
    "os"
    "log"
    "text/template"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("public")})))
  http.ListenAndServe(":8080", nil)
}

type noDir struct {
  http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
  f, err := d.Dir.Open(name)
  if err != nil {
    return nil, err
  }
  stat, err := f.Stat()
  if err != nil {
    return nil, err
  }
  if stat.IsDir() {
    return nil, os.ErrNotExist
  }
  return f, nil
}

type indexData struct {
  Name string
  List []string
}

var t = template.Must(template.ParseFiles("./template/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8 ")

  data := indexData{
    Name: "Chun Rapeepat",
    List: []string{
      "Steve",
      "Tauhoo",
      "Something",
    },
  }

  err := t.Execute(w, data)

  if err != nil {
    log.Println(err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
