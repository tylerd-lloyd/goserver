package utils_test

import (
	"restserverfd/models"
	"restserverfd/utils"
	"testing"
)

func TestFilterMetadataMap(t *testing.T) {
	testMap := make(map[int]models.Metadata)
	title1 := "hello world"
	title2 := "my cool app"
	testMap[1] = models.Metadata{Title: &title1}
	testMap[2] = models.Metadata{Title: &title2}

	testQueryParams := make(map[string][]string)
	testQueryParams["title"] = []string{"hello world"}

	result := utils.FilterMetadataMap(testMap, testQueryParams)

	if len(result) != 1 {
		t.Errorf("Filtered list is not the correct size. Expected: 1, Actual: %d", len(result))
	}

	if _, ok := result[1]; !ok {
		t.Errorf("Filtered list does not have expected item")
	}
}
