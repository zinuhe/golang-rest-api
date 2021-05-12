// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// go mod init github.com/gorilla/mux
// go run main.go

package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  Id        string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName  string `json:"lastname,omitempty"`
}

var people []Person

func GetPeople(response http.ResponseWriter, request *http.Request) {
  json.NewEncoder(response).Encode(people)
}

func GetPerson(response http.ResponseWriter, request *http.Request) {
  params := mux.Vars(request)
  for _, item := range people {
    if item.Id == params["id"] {
      json.NewEncoder(response).Encode(item)
      return
    }
  }
  json.NewEncoder(response).Encode(&Person{})
}

func AddPerson(response http.ResponseWriter, request *http.Request) {
  //json.NewEncoder(response).Encode(people)
}

func DeletePerson(response http.ResponseWriter, request *http.Request) {
  //json.NewEncoder(response).Encode(people)
}

func main() {
  router := mux.NewRouter()

  people = append(people, Person{Id: "1", FirstName: "Liz", LastName: "Bur"})
  people = append(people, Person{Id: "2", FirstName: "Jim", LastName: "Saa"})

  //End-Points
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people{id}", AddPerson).Methods("POST")
  router.HandleFunc("/people{id}", DeletePerson).Methods("DELETE")

  //Star web server and log
  //log.Fatal(http.ListenAndServe(":3000", router))
  http.ListenAndServe(":3000", router)
}
