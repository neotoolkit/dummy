package parse

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/faker"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/read"
)

type SpecType string

const (
	OpenAPI3 SpecType = "OpenAPI 3.x"
	GraphQL  SpecType = "GraphQL"
	Unknown  SpecType = "Unknown"
)

var ErrEmptySpecTypePath = errors.New("empty spec type path")

// SpecTypeError -.
type SpecTypeError struct {
	Path string
}

// Error -.
func (e *SpecTypeError) Error() string {
	return "specification type not implemented, path: " + e.Path
}

// SpecFileError -.
type SpecFileError struct {
	Path string
}

// Error -.
func (e *SpecFileError) Error() string {
	return e.Path + " without format"
}

// Parse -.
func Parse(path string) (api.API, error) {
	file, err := read.Read(path)
	if err != nil {
		return api.API{}, err
	}

	specType, err := GetSpecType(path)
	if err != nil {
		return api.API{}, err
	}

	switch specType {
	case OpenAPI3:
		var openapi openapi3.OpenAPI

		if err := yaml.Unmarshal(file, &openapi); err != nil {
			return api.API{}, err
		}

		f := faker.NewFaker()

		b := &openapi3.Builder{
			OpenAPI: openapi,
			Faker:   f,
		}

		return b.Build()
	case GraphQL:
		panic("Not implemented")
	}

	return api.API{}, nil
}

// GetSpecType returns specification type for path
func GetSpecType(path string) (SpecType, error) {
	if len(path) == 0 {
		return Unknown, ErrEmptySpecTypePath
	}

	var splitPath []string

	if path[0] == '.' {
		splitPath = strings.Split(path[1:], ".")
	} else {
		splitPath = strings.Split(path, ".")
	}

	if len(splitPath) == 1 {
		return Unknown, &SpecFileError{
			Path: path,
		}
	}

	file, err := read.Read(path)
	if err != nil {
		return Unknown, err
	}

	switch splitPath[1] {
	case "yml", "yaml":
		if err := yaml.Unmarshal(file, &openapi3.OpenAPI{}); err == nil {
			return OpenAPI3, nil
		}

		return Unknown, &SpecTypeError{
			Path: path,
		}
	case "graphql":
		return GraphQL, nil
	default:
		return Unknown, &SpecTypeError{
			Path: path,
		}
	}
}
