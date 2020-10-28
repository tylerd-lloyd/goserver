package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"restserverfd/models"
	"restserverfd/utils"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var data []models.Metadata
var dataMap map[int]models.Metadata
var counter = 1

// Run main http server
func Run() {
	dataMap = make(map[int]models.Metadata)
	router := mux.NewRouter()
	fmt.Println("Running server at localhost:50001")
	router.HandleFunc("/api/metadata", createMetadata).Methods("POST")
	router.HandleFunc("/api/metadata", searchMetadata).Methods("GET")
	router.HandleFunc("/api/metadata/{id}", getMetadata).Methods("GET")
	log.Fatal(http.ListenAndServe(":50001", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the server\n\n")
}

func createMetadata(w http.ResponseWriter, r *http.Request) {
	t := models.Metadata{}
	body, _ := ioutil.ReadAll(r.Body)
	err := yaml.UnmarshalStrict(body, &t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	if err := utils.ValidateMetadata(t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ID = counter
	counter++
	dataMap[t.ID] = t
	yaml.NewEncoder(w).Encode(t)
}

func getMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error %v", err)
		return
	}

	if m, ok := dataMap[key]; ok {
		m.ID = 0 // exclude ID field from the response object
		yaml.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "")
}

func searchMetadata(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	for p, v := range params {
		fmt.Printf("%v, %v", p, v)
	}
	var sb strings.Builder

	sb.WriteString("items:\n")
	itemSize := 0
	for _, v := range dataMap {
		itemSize++
		bytes, _ := yaml.Marshal(v)
		s := string(bytes)

		// set each item at the proper indentation level
		s = strings.ReplaceAll(s, "\n", "\n    ")
		s = strings.TrimRight(s, "\n    ")
		s = s + "\n"
		sb.WriteString(fmt.Sprintf("  -\n    %s", s))
	}

	sb.WriteString(fmt.Sprintf("count: %d\n", itemSize))

	fmt.Fprint(w, sb.String())
}
