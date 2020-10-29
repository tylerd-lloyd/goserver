package dataservice_test

import (
	"testing"

	"github.com/tylerd-lloyd/yamlserver/dataservice"
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

func TestEmailSearch(t *testing.T) {
	testMap := make(map[int]dataservice.Metadata)
	title1 := "hello world"
	title2 := "go server project"
	email1 := "john@doe.com"

	testMaintainer := dataservice.Maintainer{Email: &email1}
	testMap[1] = dataservice.Metadata{Title: &title1, Maintainers: []dataservice.Maintainer{testMaintainer}}
	testMap[2] = dataservice.Metadata{Title: &title2}

	testQueryParams := make(map[string][]string)
	testQueryParams["maintainer.email"] = []string{"john@doe.com"}

	result := dataservice.MetadataByQueryParams(testMap, testQueryParams)

	if len(result) != 1 {
		t.Errorf("Filtered list is not the correct size. Expected: 1, Actual: %d", len(result))
	}

	if _, ok := result[1]; !ok {
		t.Errorf("Filtered list does not have expected item")
	}
}
