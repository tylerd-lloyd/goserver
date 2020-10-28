package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var data []int
var counter = 0

type T struct {
	A string `yaml:"a,omitempty"`
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:"d"`
	}
}

// Run main http server
func Run() {
	router := mux.NewRouter()
	fmt.Println("Running server at localhost:50001")
	router.HandleFunc("/api/metadata", createMetadata).Methods("POST")
	router.HandleFunc("/api/metadata", metadataAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":50001", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the server\n\n")
}

func createMetadata(w http.ResponseWriter, r *http.Request) {
	t := T{}
	body, _ := ioutil.ReadAll(r.Body)
	err := yaml.UnmarshalStrict(body, &t)
	if err != nil || t.A == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err)
	} else {
		fmt.Fprintf(w, "my json payload %v", t)
	}
	data = append(data, counter)
}

func metadataAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my json payload %v", data)
}
