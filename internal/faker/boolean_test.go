package faker_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestBooleanBoolean(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker().Boolean()

	require.Equal(t, "bool", reflect.TypeOf(f.Boolean()).String())
}
