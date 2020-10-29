package dataservice_test

import (
	"gopkg.in/yamlserver/dataservice"
	"testing"
)

func TestMetadataByQueryParams(t *testing.T) {
	testMap := make(map[int]dataservice.Metadata)
	title1 := "hello world"
	title2 := "go server project"
	testMap[1] = dataservice.Metadata{Title: &title1}
	testMap[2] = dataservice.Metadata{Title: &title2}

	testQueryParams := make(map[string][]string)
	testQueryParams["title"] = []string{"hello world"}

	result := dataservice.MetadataByQueryParams(testMap, testQueryParams)

	if len(result) != 1 {
		t.Errorf("Filtered list is not the correct size. Expected: 1, Actual: %d", len(result))
	}

	if _, ok := result[1]; !ok {
		t.Errorf("Filtered list does not have expected item")
	}
}
