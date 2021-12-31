package faker_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/faker"
)

func TestPerson_FirstName(t *testing.T) {
	f := faker.NewFaker()
	firstName := f.Person().FirstName()

	require.True(t, len(firstName) > 0)
}

func TestPerson_LastName(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	lastName := p.LastName()

	require.True(t, len(lastName) > 0)
}

func TestPerson_FirstNameMale(t *testing.T) {
	f := faker.NewFaker()
	firstNameMale := f.Person().FirstNameMale()

	require.True(t, len(firstNameMale) > 0)
}

func TestPerson_FirstNameFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	firstNameFemale := p.FirstNameFemale()

	require.True(t, len(firstNameFemale) > 0)
}

func TestPerson_Name(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	name := p.Name()

	require.True(t, len(name) > 0)
	require.False(t, strings.Contains(name, "{{FirstNameMale}}"))
	require.False(t, strings.Contains(name, "{{FirstNameFemale}}"))
	require.False(t, strings.Contains(name, "{{LastName}}"))
}

func TestPerson_NameMale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	nameMale := p.NameMale()

	require.True(t, len(nameMale) > 0)
}

func TestPerson_NameFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	nameFemale := p.NameFemale()

	require.True(t, len(nameFemale) > 0)
}

func TestPerson_Gender(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	gender := p.Gender()

	require.True(t, gender == "Male" || gender == "Female")
}

func TestPerson_GenderMale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	genderMale := p.GenderMale()

	require.True(t, genderMale == "Male")
}

func TestPerson_GenderFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	genderFemale := p.GenderFemale()

	require.True(t, genderFemale == "Female")
}
