package validators_test

import (
	"gopkg.in/yamlserver/dataservice"
	"gopkg.in/yamlserver/validators"
	"testing"

	"gopkg.in/yaml.v2"
)

var testData = `
title: hello
version: 1.2.3
maintainers:
  - name: Tyler
    email: ty@email.com
company: msft
website: golang.org
source: blank.com
license: gnu
description: |
  my multiline description
  which supports markdown`

func TestValidMetadata_ShouldNotReturnError(t *testing.T) {
	testM := dataservice.Metadata{}
	yaml.Unmarshal([]byte(testData), &testM)
	err := validators.ValidateMetadata(testM)
	if err != nil {
		t.Error(err)
	}
}

func TestValidMetadata_ShouldReturnError(t *testing.T) {
	testM := dataservice.Metadata{}
	yaml.Unmarshal([]byte(testData), &testM)
	testM.Company = nil
	testM.Source = nil

	err := validators.ValidateMetadata(testM)
	if err == nil {
		t.Error("Expected error, no error returned")
	}
}
