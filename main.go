// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// go run main.go

package main

import (
  "fmt"
  "net/http"
)

type Person struct {
  Id        string `json:"id,omitempty"`
  FirstName string `json:"firstname,omitempty"`
  LastName  string `json:"lastname,omitempty"`
}

var people []Person

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "homePage EndPoint")
  fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
  http.HandleFunc("/", homePage)

  //Star web server and log
  //log.Fatal(http.ListenAndServe(":3000", router))
  http.ListenAndServe(":1000", nil)
}

func main() {
  handleRequests()
}
