package openapi3

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func Parse(path string) (res OpenAPI, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(file, &res); err != nil {
		return
	}

	return
}
