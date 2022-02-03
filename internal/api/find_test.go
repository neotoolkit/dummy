package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/dummy/internal/api"
)

type ResponseParamsBody struct {
	Body         interface{}
	IsBodyBroken bool
}

func (params ResponseParamsBody) Read(p []byte) (n int, err error) {
	jsonData, err := json.Marshal(params.Body)
	if params.IsBodyBroken || err != nil {
		jsonData = []byte{1, 2, 3}
	}

	readerResult, readerError := bytes.NewReader(jsonData).Read(p)

	return readerResult, readerError
}

func (ResponseParamsBody) Close() error {
	return nil
}

func TestIsPathMatchTemplate(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		param string
		want  bool
	}{
		{
			name:  "",
			path:  "",
			param: "",
			want:  true,
		},
		{
			name:  "",
			path:  "/path",
			param: "/path",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1",
			param: "/path/{1}",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1/path/2",
			param: "/path/{1}/path/{2}",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1/path1/2",
			param: "/path/{1}/path/{2}",
			want:  false,
		},
		{
			name:  "",
			path:  "/path/1/path/2",
			param: "/path/{1}/path1/{2}",
			want:  false,
		},
		{
			name:  "",
			path:  "/path/1/path/2/path",
			param: "/path/{1}/path/{2}",
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := api.IsPathMatchTemplate(tc.path, tc.param)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestFindResponseError(t *testing.T) {
	got := &api.FindResponseError{
		Method: "test method",
		Path:   "test path",
	}

	require.Equal(t, got.Error(), "not specified operation: test method test path")
}

func TestFindResponse(t *testing.T) {
	bodyFixedPath := map[string]api.FieldType{
		"param1": {
			Required: true,
			Type:     "string",
		},
		"param2": {
			Required: false,
			Type:     "string",
		},
	}

	responseFixedPathJSON := api.Response{
		StatusCode: 200,
		MediaType:  "application/json",
	}
	responseFixedPathZip := api.Response{
		StatusCode: 200,
		MediaType:  "application/zip",
	}

	operationFixedPath := api.Operation{
		Method: "POST",
		Path:   "some/fixed/path",
		Body:   bodyFixedPath,
		Responses: []api.Response{
			responseFixedPathJSON,
			responseFixedPathZip,
		},
	}

	a := api.API{
		Operations: []api.Operation{
			operationFixedPath,
		},
	}

	mismatchByOperationPathError := api.FindResponseError{
		Method: "POST",
		Path:   "some/other/path",
	}
	mismatchByOperationMethodError := api.FindResponseError{
		Method: "GET",
		Path:   "some/fixed/path",
	}
	bodyWithoutRequiredParamError := errors.New("empty require field")

	bodyWithoutRequiredParam := map[string]interface{}{
		"param3": "qwe",
		"param2": "rty",
	}
	bodyWithoutOptionalParam := map[string]interface{}{
		"param3": "qwe",
		"param1": "rty",
	}
	bodyWithAllParam := map[string]interface{}{
		"param1": "qwe",
		"param2": "rty",
	}

	emptyResponse := api.Response{}

	tests := []struct {
		name              string
		path              string
		method            string
		body              interface{}
		wantFirst         api.Response
		wantSecond        interface{}
		responseMediaType string
	}{
		{
			name:       "Mismatch by operation path",
			path:       "some/other/path",
			method:     "POST",
			body:       bodyFixedPath,
			wantFirst:  emptyResponse,
			wantSecond: &mismatchByOperationPathError,
		},
		{
			name:       "Mismatch by operation method",
			path:       "some/fixed/path",
			method:     "GET",
			body:       bodyFixedPath,
			wantFirst:  emptyResponse,
			wantSecond: &mismatchByOperationMethodError,
		},
		{
			name:       "Mismatch by body without required params",
			path:       "some/fixed/path",
			method:     "POST",
			body:       bodyWithoutRequiredParam,
			wantFirst:  emptyResponse,
			wantSecond: bodyWithoutRequiredParamError,
		},
		{
			name:       "Match by body without optional params",
			path:       "some/fixed/path",
			method:     "POST",
			body:       bodyWithoutOptionalParam,
			wantFirst:  responseFixedPathJSON,
			wantSecond: nil,
		},
		{
			name:       "Match by response without selected media type",
			path:       "some/fixed/path",
			method:     "POST",
			body:       bodyWithAllParam,
			wantFirst:  responseFixedPathJSON,
			wantSecond: nil,
		},
		{
			name:              "Match by all criteria",
			path:              "some/fixed/path",
			method:            "POST",
			body:              bodyWithAllParam,
			wantFirst:         responseFixedPathZip,
			wantSecond:        nil,
			responseMediaType: "application/zip",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bodyReader := ResponseParamsBody{Body: tc.body}
			mediaType := "application/xml"

			if tc.responseMediaType != "" {
				mediaType = tc.responseMediaType
			}
			params := api.FindResponseParams{
				Path:      tc.path,
				Method:    tc.method,
				Body:      bodyReader,
				MediaType: mediaType,
			}
			firstResult, secondResult := a.FindResponse(params)

			require.Equal(t, tc.wantFirst, firstResult)
			require.Equal(t, tc.wantSecond, secondResult)
		})
	}

	t.Run("Broken json in body reader", func(t *testing.T) {
		params := api.FindResponseParams{
			Path:      "some/fixed/path",
			Method:    "POST",
			Body:      ResponseParamsBody{Body: bodyWithAllParam, IsBodyBroken: true},
			MediaType: "application/xml",
		}

		_, err := a.FindResponse(params)
		require.Error(t, err)
	})
}
