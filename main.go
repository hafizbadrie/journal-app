package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func Index(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
  fmt.Fprint(writer, "Hello World!")
}

func main() {
  fmt.Println("Preparing webserver...")
  router := httprouter.New()
  router.GET("/", Index)

  http.ListenAndServe(":8080", router)
}
