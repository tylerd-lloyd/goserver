package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy
b:
  c: 2
  d: [3,4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("Running server at localhost:50001")
	http.ListenAndServe(":50001", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Fprintf(w, "Welcome! %v", t)
}
