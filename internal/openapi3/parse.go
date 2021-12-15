package openapi3

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-yaml/yaml"
)

func Parse(path string) (*OpenAPI, error) {
	file, err := read(path)
	if err != nil {
		return nil, err
	}

	var res OpenAPI

	if err := yaml.Unmarshal(file, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func read(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return readFromURL(path)
	}

	return readFromFile(path)
}

func readFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func readFromFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
