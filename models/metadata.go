package models

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
