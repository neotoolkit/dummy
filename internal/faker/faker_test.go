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

func TestByName(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		faker string
	}

	tests := []test{
		{
			name:  "",
			faker: "Boolean",
		},
		{
			name:  "",
			faker: "Username",
		},
		{
			name:  "",
			faker: "GTLD",
		},
		{
			name:  "",
			faker: "Domain",
		},
		{
			name:  "",
			faker: "Email",
		},
		{
			name:  "",
			faker: "Person.FirstName",
		},
		{
			name:  "",
			faker: "Person.LastName",
		},
		{
			name:  "",
			faker: "Person.FirstNameMale",
		},
		{
			name:  "",
			faker: "Person.FirstNameFemale",
		},
		{
			name:  "",
			faker: "Person.Name",
		},
		{
			name:  "",
			faker: "Person.NameMale",
		},
		{
			name:  "",
			faker: "Person.NameFemale",
		},
		{
			name:  "",
			faker: "Person.Gender",
		},
		{
			name:  "",
			faker: "Person.GenderMale",
		},
		{
			name:  "",
			faker: "Person.GenderFemale",
		},
		{
			name:  "",
			faker: "UUID",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := faker.NewFaker()
			got := f.ByName(tc.faker)

			require.True(t, got != nil)
		})
	}
}
