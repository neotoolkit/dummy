package faker_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestUsername(t *testing.T) {
	i := faker.NewFaker().Internet()
	username := i.Username()

	require.Equal(t, true, len(username) > 0)
	require.Equal(t, false, strings.Contains(username, " "))
}

func TestGTLD(t *testing.T) {
	i := faker.NewFaker().Internet()
	gTLD := i.GTLD()

	require.Equal(t, true, len(gTLD) > 0)
}

func TestDomain(t *testing.T) {
	i := faker.NewFaker().Internet()
	d := i.Domain()

	require.Equal(t, true, len(d) > 0)
}

func TestEmail(t *testing.T) {
	i := faker.NewFaker().Internet()
	e := i.Email()

	require.Equal(t, true, len(e) > 0)
}
