package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  handlers "github.com/hafizbadrie/journal-app/handlers"
)

func Index(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  //username, password, _ := req.BasicAuth()

  //if username != httpUsername || password != httpPassword {
    //http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
    //return &errorAuth{"Unauthorized Access"}
  //}

  fmt.Fprint(writer, "Hello World!")


  //return nil
}

func main() {
  fmt.Println("Preparing webserver to run on :8080...")
  router := httprouter.New()
  //router.GET("/", HTTP(Index))
  router.GET("/", Index)
  router.GET("/employee", Show)
  router.GET("/journals", handlers.Journals)

  http.ListenAndServe(":8080", router)
}
