package parse

import (
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

// SpecTypeError -.
type SpecTypeError struct {
	Path string
}

// Error -.
func (e *SpecTypeError) Error() string {
	return "specification type not implemented, path: " + e.Path
}

// Parse -.
func Parse(path string) (api.API, error) {
	file, err := read.Read(path)
	if err != nil {
		return api.API{}, err
	}

	specType, err := specType(path)
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

// specType returns specification type for path
func specType(path string) (SpecType, error) {
	splitPath := strings.Split(path, ".")

	switch splitPath[1] {
	case "yml", "yaml":
		return OpenAPI3, nil
	case "graphql":
		return GraphQL, nil
	default:
		return Unknown, &SpecTypeError{
			Path: path,
		}
	}
}
