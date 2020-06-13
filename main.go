package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  handlers "github.com/hafizbadrie/journal-app/handlers"
)

// TESTING: json and functions
type EmployeeData struct {
  Data []Employee `json:"data"`
  Status string `json:"status"`
}

type Employee struct {
  FirstName string `json:"first_name"`
}

func (emp Employee) getFirstName() string {
  return emp.FirstName
}

func (emp *Employee) changeFirstName(firstName string) {
  emp.FirstName = firstName
}

func Show(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  var employees []Employee

  empHafiz := Employee{
    FirstName: "Hafiz",
  }
  empBadrie := Employee{
    FirstName: "Badrie",
  }
  empLubis := Employee{
    FirstName: "Lubis",
  }
  employees = append(employees, empHafiz)
  employees = append(employees, empBadrie, empLubis)
  employeeData := EmployeeData{
    Data: employees,
    Status: "OK",
  }
  var jsonData []byte
  jsonData, err := json.Marshal(employeeData)
  if err != nil  {
    panic(err)
  }

  fmt.Println(string(jsonData))
  fmt.Fprint(writer, string(jsonData))
}

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
