package validators

import (
	"fmt"
	"gopkg.in/yamlserver/dataservice"
	"regexp"
	"strings"
)

type argError struct {
	msg string
}

var validEmail = regexp.MustCompile(`^[A-z0-9._%+-]+@[A-z0-9.-]+\.[A-z]{2,}$`)

func (e *argError) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

func writeSchemaError(sb *strings.Builder, fieldName string) {
	if sb.Len() == 0 {
		sb.WriteString("DataSchemaError:\n")
	}
	sb.WriteString(fmt.Sprintf("- MissingField %s\n", fieldName))
}

// ValidateMetadata checks all fields for non-nil references and valid email address formats
func ValidateMetadata(m dataservice.Metadata) error {
	var sb strings.Builder
	if m.Company == nil {
		writeSchemaError(&sb, "Company")
	}

	if m.Description == nil {
		writeSchemaError(&sb, "Description")
	}

	if m.License == nil {
		writeSchemaError(&sb, "License")
	}

	if m.Source == nil {
		writeSchemaError(&sb, "Source")
	}

	if m.Title == nil {
		writeSchemaError(&sb, "Title")
	}

	if m.Version == nil {
		writeSchemaError(&sb, "Version")
	}

	if m.Website == nil {
		writeSchemaError(&sb, "Website")
	}

	if len(m.Maintainers) == 0 {
		writeSchemaError(&sb, "Maintainers")
	}

	for _, maintainer := range m.Maintainers {
		if maintainer.Email == nil {
			writeSchemaError(&sb, "Maintainer.Email")
		}

		if maintainer.Name == nil {
			writeSchemaError(&sb, "Maintainer.Name")
		}

		if !validEmail.MatchString(*maintainer.Email) {
			if sb.Len() == 0 {
				sb.WriteString("DataSchemaError:\n")
			}
			sb.WriteString("- Malformed email " + *maintainer.Email)
		}
	}
	if sb.Len() == 0 {
		return nil
	}

	return &argError{sb.String()}
}
