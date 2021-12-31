package faker_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestUUID_v4(t *testing.T) {
	f := faker.NewFaker()
	value := f.UUID().V4()
	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", value)

	require.NoError(t, err)
	require.True(t, match)
}

func TestUUID_V4UniqueInSequence(t *testing.T) {
	f := faker.NewFaker()
	last := f.UUID().V4()
	current := f.UUID().V4()

	require.Equal(t, true, last != current)
}
