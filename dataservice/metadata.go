package dataservice

// Metadata represents application metadata expected to be unmarshaled from yaml
type Metadata struct {
	ID          int          `yaml:",omitempty"`
	Title       *string      `yaml:"title,omitempty"`
	Version     *string      `yaml:"version,omitempty"`
	Maintainers []Maintainer `yaml:"maintainers"`
	Company     *string      `yaml:",omitempty"`
	Website     *string      `yaml:",omitempty"`
	Source      *string      `yaml:",omitempty"`
	License     *string      `yaml:",omitempty"`
	Description *string      `yaml:",omitempty"`
}

// Maintainer is a person maintaining the application
type Maintainer struct {
	Name  *string `yaml:"name,omitempty"`
	Email *string `yaml:"email,omitempty"`
}

// MetadataByQueryParams returns a filtered copy of the original data map with the items that match the search params
func MetadataByQueryParams(m map[int]Metadata, searchParams map[string][]string) map[int]Metadata {
	newMap := make(map[int]Metadata)
	for k, v := range m {
		newMap[k] = v
	}

	if p, ok := searchParams["title"]; ok {
		for k, v := range m {
			if contains(p, *v.Title) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["version"]; ok {
		for k, v := range m {
			if contains(p, *v.Version) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["company"]; ok {
		for k, v := range m {
			if contains(p, *v.Company) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["website"]; ok {
		for k, v := range m {
			if contains(p, *v.Website) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["source"]; ok {
		for k, v := range m {
			if contains(p, *v.Source) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["license"]; ok {
		for k, v := range m {
			if contains(p, *v.License) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["description"]; ok {
		for k, v := range m {
			if contains(p, *v.Description) == -1 {
				delete(newMap, k)
			}
		}
	}

	if p, ok := searchParams["maintainer.email"]; ok {
		for k, v := range m {
			if len(v.Maintainers) == 0 {
				delete(newMap, k)
			}
			for _, maintainer := range v.Maintainers {
				if contains(p, *maintainer.Email) == -1 {
					delete(newMap, k)
				}
			}
		}
	}

	if p, ok := searchParams["maintainer.name"]; ok {
		for k, v := range m {
			if len(v.Maintainers) == 0 {
				delete(newMap, k)
			}
			for _, maintainer := range v.Maintainers {
				if contains(p, *maintainer.Name) == -1 {
					delete(newMap, k)
				}
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
