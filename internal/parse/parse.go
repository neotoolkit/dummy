package parse

import (
	"errors"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/neotoolkit/faker"
	"github.com/neotoolkit/openapi"

	"github.com/neotoolkit/dummy/internal/api"
	"github.com/neotoolkit/dummy/internal/read"
)

type SpecType string

var Test = 2

const (
	OpenAPI SpecType = "OpenAPI"
	GraphQL SpecType = "GraphQL"
	Unknown SpecType = "Unknown"
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
	case OpenAPI:
		oapi, err := openapi.Parse(file)
		if err != nil {
			return api.API{}, err
		}

		f := faker.NewFaker()

		b := &api.Builder{
			OpenAPI: oapi,
			Faker:   f,
		}

		return b.Build()
	case GraphQL:
		return api.API{}, nil
	}

	return api.API{}, nil
}

// GetSpecType returns specification type for path
func GetSpecType(path string) (SpecType, error) {
	if len(path) == 0 {
		return Unknown, ErrEmptySpecTypePath
	}

	splitPath := strings.Split(path[1:], ".")

	if len(splitPath) == 1 {
		return Unknown, &SpecFileError{
			Path: path,
		}
	}

	file, err := read.Read(path)
	if err != nil {
		return Unknown, err
	}

	switch splitPath[len(splitPath)-1] {
	case "yml", "yaml":
		var openapi openapi.OpenAPI

		err := yaml.Unmarshal(file, &openapi)
		if err != nil || len(openapi.OpenAPI) == 0 {
			return Unknown, &SpecTypeError{
				Path: path,
			}
		}

		return OpenAPI, nil
	case "graphql":
		return GraphQL, nil
	default:
		return Unknown, &SpecTypeError{
			Path: path,
		}
	}
}
