package faker_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestBoolean_Boolean(t *testing.T) {
	f := faker.NewFaker().Boolean()

	require.Equal(t, "bool", reflect.TypeOf(f.Boolean()).String())
}
