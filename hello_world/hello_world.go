package main

import (
  "fmt"
  "net/http"
)

func main(){
  fmt.Println("Server inicializado")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola mundo en esta mierda %s", r.URL.Path)
  })

  http.ListenAndServe(":80", nil)
}