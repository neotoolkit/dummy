package faker_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestBooleanBoolean(t *testing.T) {
	f := faker.NewFaker().Boolean()

	require.True(t, f.Boolean() == true || f.Boolean() == false)
}
