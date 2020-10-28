package main

import (
	"fmt"
	"net/http"

	"restserverfd/server"

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
	server.Run()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t := T{}

	err := yaml.UnmarshalStrict([]byte(data), &t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err)
	} else {
		fmt.Fprintf(w, "Welcome! %v", t)
	}
}
