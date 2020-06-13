package handlers

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
  "net/http"
  "net/http/httptest"
  "github.com/julienschmidt/httprouter"
  "testing"
  "github.com/stretchr/testify/assert"
  "encoding/json"
)

// Test cases
// 1. [OK] Given endpoint /journals, it will give empty data when data is non existent
// 2. [OK] Given endpoint /journals, it will give data from the database
// 3. Given endpoint /journals, it will give data from the database limit by 50 journals
// 4. Given endpoint /journals and page, it will take data from the next page
// 5. Given endpoint /journals and page, it will give empty result if page is more than the limit
// 6. Given endpoint /journals and page, it will return 400 when page is not a valid number

// Scenario: should return empty result when data is also empty
func TestJournalsEndpointWithEmptyData(t *testing.T) {
  httpRecorder := httptest.NewRecorder()
  router := httprouter.New()
  router.GET("/journals", Journals)
  req, err := http.NewRequest("GET", "/journals", nil)
  if err != nil {
    t.Fatal(err)
  }
  router.ServeHTTP(httpRecorder, req)

  var journalData []JournalData
  responseCode := httpRecorder.Code
  responseBody := httpRecorder.Body.String()
  expectedRawBody := JournalResponse{
    journalData,
    "OK",
  }
  expectedBody, _ := json.Marshal(expectedRawBody)
  assert.Equal(t, 200, responseCode, "Response code should be 200")
  assert.Equal(t, string(expectedBody), responseBody, "Response should be match")
}

// Scenario: should return some data
func TestJournalsEndpointWithData(t *testing.T) {
  // Insert data to PSQL
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)
  db, err := sql.Open("postgres", psqlInfo)
  defer db.Close()
  if err != nil {
    panic(err)
  }
  sqlStatement := `
  INSERT INTO journals (id, name, created_at, updated_at)
  VALUES 
    (1, 'My First Journal', '2020-05-02 00:00:00',  '2020-05-02 00:00:00'),
    (2, 'My Second Journal', '2020-05-03 00:00:00', '2020-05-03 00:00:00'),
    (3, 'My Third Journal', '2020-05-04 00:00:00', '2020-05-04 00:00:00')
  `
  _, err = db.Exec(sqlStatement)
  if err != nil {
    panic(err)
  }

  // Test codes here
  httpRecorder := httptest.NewRecorder()
  router := httprouter.New()
  router.GET("/journals", Journals)
  req, err := http.NewRequest("GET", "/journals", nil)
  if err != nil {
    t.Fatal(err)
  }
  router.ServeHTTP(httpRecorder, req)

  var journalData []JournalData
  journalData = append(journalData, JournalData{1, "My First Journal", "2020-05-02T00:00:00Z", "2020-05-02T00:00:00Z"})
  journalData = append(journalData, JournalData{2, "My Second Journal", "2020-05-03T00:00:00Z", "2020-05-03T00:00:00Z"})
  journalData = append(journalData, JournalData{3, "My Third Journal", "2020-05-04T00:00:00Z", "2020-05-04T00:00:00Z"})
  responseCode := httpRecorder.Code
  responseBody := httpRecorder.Body.String()
  expectedRawBody := JournalResponse{
    journalData,
    "OK",
  }
  expectedBody, _ := json.Marshal(expectedRawBody)
  assert.Equal(t, 200, responseCode, "Response code should be 200")
  assert.Equal(t, string(expectedBody), responseBody, "Response should be match")

  // Clean up  PSQL
  sqlStatement = `
  TRUNCATE journals
  `
  _,  err = db.Exec(sqlStatement)
  if err != nil {
    panic(err)
  }
}

