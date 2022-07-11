package parse

import (
	"errors"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/neotoolkit/openapi"

	"github.com/neotoolkit/dummy/internal/api"
	"github.com/neotoolkit/dummy/internal/read"
)

type SpecType string

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

		b := &api.Builder{
			OpenAPI: oapi,
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

	p := filepath.Ext(path)

	if len(p) == 0 {
		return Unknown, &SpecFileError{Path: path}
	}

	file, err := read.Read(path)
	if err != nil {
		return Unknown, err
	}

	switch p {
	case ".yml", ".yaml":
		var oapi openapi.OpenAPI

		if err := yaml.Unmarshal(file, &oapi); err != nil || len(oapi.OpenAPI) == 0 {
			return Unknown, &SpecTypeError{Path: path}
		}

		return OpenAPI, nil
	case ".graphql":
		return GraphQL, nil
	default:
		return Unknown, &SpecFileError{Path: path}
	}
}
