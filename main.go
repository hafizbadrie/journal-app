package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"

  "database/sql"
  _ "github.com/lib/pq"
)

const (
  host   = "localhost"
  port   = 5432
  user   = "postgres"
  dbname = "redash"
)

func Index(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  fmt.Fprint(writer, "Hello World!")
}

func Sum(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println("Connecting to database...")
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  fmt.Println("Preparing webserver to run on :8080...")
  router := httprouter.New()
  router.GET("/", Index)

  http.ListenAndServe(":8080", router)
}
