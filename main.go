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
  httpUsername = "hafizbadrie"
  httpPassword = "1234567890"
)

type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error
type errorAuth struct {
  msg string
}

func (e *errorAuth) Error() string {
  return e.msg
}

func HTTP(handle HandleWithError) httprouter.Handle {
    return func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
      handle(writer, req, params)
    }
}

func Index(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) error {
  username, password, _ := req.BasicAuth()

  if username != httpUsername || password != httpPassword {
    http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
    return &errorAuth{"Unauthorized Access"}
  }

  fmt.Fprint(writer, "Hello World!")

  return nil
}

// This is a sample function to add unit test
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
  router.GET("/", HTTP(Index))

  http.ListenAndServe(":8080", router)
}
