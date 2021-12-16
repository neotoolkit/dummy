package openapi3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	userSchema = Schema{
		Properties: Schemas{
			"id": &Schema{
				Type:    "string",
				Format:  "uuid",
				Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
			},
			"firstName": &Schema{
				Type:    "string",
				Example: "Larry",
			},
			"lastName": &Schema{
				Type:    "string",
				Example: "Page",
			},
		},
		Type: "object",
	}

	uuidSchema = Schema{
		Type:    "string",
		Format:  "uuid",
		Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
	}
)

type schemaContextStub struct{}

func (s schemaContextStub) LookupByReference(ref string) (Schema, error) {
	switch ref {
	case "#/components/schemas/User":
		return userSchema, nil
	case "#/components/schemas/UUID":
		return uuidSchema, nil
	default:
		return Schema{}, fmt.Errorf("unknown schema: %q", ref)
	}
}

func TestSchema_ResponseByExample(t *testing.T) {
	type fields struct {
		Properties Schemas
		Type       string
		Format     string
		Default    interface{}
		Example    interface{}
		Faker      string
		Items      *Schema
		Reference  string
	}

	type args struct {
		schemaContext schemaContext
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes interface{}
		wantErr bool
	}{
		{
			name: "Simple schema",
			fields: fields{
				Properties: Schemas{
					"id": &Schema{
						Type:    "string",
						Format:  "uuid",
						Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
					},
					"firstName": &Schema{
						Type:    "string",
						Example: "Larry",
					},
					"lastName": &Schema{
						Type:    "string",
						Example: "Page",
					},
				},
				Type: "object",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
		{
			name: "Simple schema with reference",
			fields: fields{
				Reference: "#/components/schemas/User",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
		{
			name: "Array schema with reference",
			fields: fields{
				Type: "array",
				Items: &Schema{
					Reference: "#/components/schemas/User",
				},
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: []interface{}{
				map[string]interface{}{
					"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
					"firstName": "Larry",
					"lastName":  "Page",
				},
			},
			wantErr: false,
		},
		{
			name: "Schema property with reference",
			fields: fields{
				Properties: Schemas{
					"id": &Schema{
						Reference: "#/components/schemas/UUID",
					},
					"firstName": &Schema{
						Type:    "string",
						Example: "Larry",
					},
					"lastName": &Schema{
						Type:    "string",
						Example: "Page",
					},
				},
				Type: "object",
			},
			args: args{
				schemaContext: schemaContextStub{},
			},
			wantRes: map[string]interface{}{
				"id":        "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a",
				"firstName": "Larry",
				"lastName":  "Page",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Schema{
				Properties: tt.fields.Properties,
				Type:       tt.fields.Type,
				Format:     tt.fields.Format,
				Default:    tt.fields.Default,
				Example:    tt.fields.Example,
				Faker:      tt.fields.Faker,
				Items:      tt.fields.Items,
				Reference:  tt.fields.Reference,
			}
			gotRes, err := s.ResponseByExample(tt.args.schemaContext)

			require.NoError(t, err)
			require.Equal(t, tt.wantRes, gotRes)
		})
	}
}
