package utils

import "restserverfd/models"

func FilterMetadataMap(m map[int]models.Metadata, searchParams map[string][]string) map[int]models.Metadata {
	newMap := make(map[int]models.Metadata)
	for k, v := range m {
		newMap[k] = v
	}

	if f, ok := searchParams["title"]; ok {
		for k, v := range m {
			if contains(f, *v.Title) == -1 {
				delete(newMap, k)
			}
		}
	}

	return newMap
}

func contains(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
