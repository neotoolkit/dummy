package faker_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestIntBetween(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	value := f.IntBetween(1, 100)

	require.Equal(t, fmt.Sprintf("%T", value), "int")
	require.True(t, value >= 1)
	require.True(t, value <= 100)
}
