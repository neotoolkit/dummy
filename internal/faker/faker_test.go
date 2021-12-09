package faker_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestIntBetween(t *testing.T) {
	f := faker.NewFaker()
	value := f.IntBetween(1, 100)

	require.Equal(t, fmt.Sprintf("%T", value), "int")
	require.Equal(t, true, value >= 1)
	require.Equal(t, true, value <= 100)
}
