package handlers

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"

  "database/sql"
  _ "github.com/lib/pq"
)

const (
  dbHost = "localhost"
  dbPort = 5432
  dbUser = "hafizbadrielubis"
  dbName = "journal_app_development"

  httpUsername = "hafizbadrie"
  httpPassword = "1234567890"
)

type JournalResponse struct {
  Data []JournalData `json:"data"`
  Status string `json:"status"`
}

type JournalData struct {
  Id int `json:"id"`
  Name string `json:"name"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}

func Journals(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  // Create connection
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)
  db, err := sql.Open("postgres", psqlInfo)
  defer db.Close()
  if err != nil {
    panic(err)
  }

  // Getting data and iterating the data
  sqlStatement := "SELECT * FROM journals"
  rows, err := db.Query(sqlStatement)
  if err != nil {
    panic(err)
  }

  var data []JournalData
  for rows.Next() {
    var journalData JournalData

    err := rows.Scan(&journalData.Id, &journalData.Name, &journalData.CreatedAt, &journalData.UpdatedAt)
    if err != nil {
      panic(err)
    }

    data = append(data, journalData)
  }

  // Preparing the response
  jsonData := JournalResponse{
    data,
    "OK",
  }
  response, err := json.Marshal(jsonData)
  if err != nil {
    panic(err)
  }

  // Write the response out
  fmt.Fprint(writer, string(response))
}
