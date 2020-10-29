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
func Run(port string) {
	dataMap = make(map[int]models.Metadata, 100)
	router := mux.NewRouter()
	router.HandleFunc("/api/metadata", createMetadata).Methods("POST")
	router.HandleFunc("/api/metadata", searchMetadata).Methods("GET")
	router.HandleFunc("/api/metadata/{id}", getMetadata).Methods("GET")
	fmt.Printf("Server running at localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
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
		log.Printf("BadRequest 400 - %v\n", err)
		return
	}

	if err := utils.ValidateMetadata(t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		log.Printf("BadRequest 400 - %s\n", err.Error())
		return
	}
	t.ID = counter
	counter++
	dataMap[t.ID] = t
	w.WriteHeader(http.StatusCreated)
	yaml.NewEncoder(w).Encode(t)
	log.Printf("Created 201 - OK\n")
}

func getMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error %v", err)
		log.Printf("BadRequest 400 - %v\n", err)
		return
	}

	if m, ok := dataMap[key]; ok {
		m.ID = 0 // exclude ID field from the response object
		yaml.NewEncoder(w).Encode(m)
		log.Printf("Success 200 - OK\n")
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "")
	log.Printf("NotFound 404 - id=%d\n", key)
}

func searchMetadata(w http.ResponseWriter, r *http.Request) {
	params := make(map[string][]string)

	for p, v := range r.URL.Query() {
		params[strings.ToLower(p)] = v
	}

	filtered := utils.FilterMetadataMap(dataMap, params)

	fmt.Fprint(w, buildResponse(filtered))
	log.Printf("Success 200 - OK\n")
}

func buildResponse(data map[int]models.Metadata) string {
	var itemSize int
	var sb strings.Builder
	sb.WriteString("items:\n")
	for _, v := range data {
		itemSize++
		bytes, _ := yaml.Marshal(v)
		s := string(bytes)

		// set each item at the proper indentation level
		s = strings.ReplaceAll(s, "\n", "\n    ")
		s = strings.TrimRight(s, "\n    ")
		s = s + "\n"
		sb.WriteString(fmt.Sprintf("  -\n    %s", s))
		if itemSize == 100 {
			break // limit response size to 100 items
		}
	}

	sb.WriteString(fmt.Sprintf("count: %d\n", itemSize))

	return sb.String()
}
