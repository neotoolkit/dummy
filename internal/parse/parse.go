package parse

import (
	"gopkg.in/yaml.v3"

	"github.com/go-dummy/dummy/internal/openapi3"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/faker"
	"github.com/go-dummy/dummy/internal/read"
)

// Parse -.
func Parse(path string) (api.API, error) {
	file, err := read.Read(path)
	if err != nil {
		return api.API{}, err
	}

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
}
