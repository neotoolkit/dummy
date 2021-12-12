package faker_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestFirstName(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	firstName := f.Person().FirstName()

	require.True(t, len(firstName) > 0)
}

func TestLastName(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	lastName := p.LastName()

	require.True(t, len(lastName) > 0)
}

func TestFirstNameMale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	firstNameMale := f.Person().FirstNameMale()

	require.True(t, len(firstNameMale) > 0)
}

func TestFirstNameFemale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	firstNameFemale := p.FirstNameFemale()

	require.True(t, len(firstNameFemale) > 0)
}

func TestName(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	name := p.Name()

	require.True(t, len(name) > 0)
	require.False(t, strings.Contains(name, "{{FirstNameMale}}"))
	require.False(t, strings.Contains(name, "{{FirstNameFemale}}"))
	require.False(t, strings.Contains(name, "{{LastName}}"))
}

func TestNameMale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	nameMale := p.NameMale()

	require.True(t, len(nameMale) > 0)
}

func TestNameFemale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	nameFemale := p.NameFemale()

	require.True(t, len(nameFemale) > 0)
}

func TestGender(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	gender := p.Gender()

	require.True(t, gender == "Male" || gender == "Female")
}

func TestGenderMale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	genderMale := p.GenderMale()

	require.True(t, genderMale == "Male")
}

func TestGenderFemale(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()
	p := f.Person()
	genderFemale := p.GenderFemale()

	require.True(t, genderFemale == "Female")
}
