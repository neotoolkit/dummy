package model

import (
	"io"
)

type Response interface {
	StatusCode() int
	ExampleValue(exampleKey string) interface{}
}

// FindResponseParams -.
type FindResponseParams struct {
	Path      string
	Method    string
	Body      io.ReadCloser
	MediaType string
}

type API interface {
	FindResponse(params FindResponseParams) (Response, error)
}
