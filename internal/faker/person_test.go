package faker_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestFirstName(t *testing.T) {
	f := faker.NewFaker()
	firstName := f.Person().FirstName()

	require.True(t, len(firstName) > 0)
}

func TestLastName(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	lastName := p.LastName()

	require.True(t, len(lastName) > 0)
}
