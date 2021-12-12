package faker

import (
	"strings"
)

// Person is struct for Person
type Person struct {
	Faker *Faker
}

// FirstName returns random first name
func (p Person) FirstName() string {
	firstName := make([]string, len(p.Faker.firstNameMale)+len(p.Faker.firstNameFemale))
	i := 0

	for j := 0; j < len(p.Faker.firstNameMale); j++ {
		firstName[i] = p.Faker.firstNameMale[j]
		i++
	}

	for j := 0; j < len(p.Faker.firstNameFemale); j++ {
		firstName[i] = p.Faker.firstNameFemale[j]
		i++
	}

	return p.Faker.RandomStringElement(firstName)
}

// LastName returns random last name
func (p Person) LastName() string {
	i := p.Faker.IntBetween(0, len(p.Faker.lastName)-1)
	return p.Faker.lastName[i]
}

// FirstNameMale returns random male first name
func (p Person) FirstNameMale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.firstNameMale)-1)
	return p.Faker.firstNameMale[i]
}

// FirstNameFemale returns random female first name
func (p Person) FirstNameFemale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.firstNameFemale)-1)
	return p.Faker.firstNameFemale[i]
}

// Name returns random name
func (p Person) Name() string {
	format := make([]string, len(p.Faker.maleNameFormat)+len(p.Faker.femaleNameFormat))
	i := 0

	for j := 0; j < len(p.Faker.maleNameFormat); j++ {
		format[i] = p.Faker.maleNameFormat[j]
		i++
	}

	for j := 0; j < len(p.Faker.femaleNameFormat); j++ {
		format[i] = p.Faker.femaleNameFormat[j]
		i++
	}

	name := format[p.Faker.IntBetween(0, len(format)-1)]

	// {{firstNameMale}}
	if strings.Contains(name, "{{firstNameMale}}") {
		name = strings.ReplaceAll(name, "{{firstNameMale}}", p.FirstNameMale())
	}

	// {{firstNameFemale}}
	if strings.Contains(name, "{{firstNameFemale}}") {
		name = strings.ReplaceAll(name, "{{firstNameFemale}}", p.FirstNameFemale())
	}

	// {{lastName}}
	if strings.Contains(name, "{{lastName}}") {
		name = strings.ReplaceAll(name, "{{lastName}}", p.LastName())
	}

	return name
}

// NameMale returns random male name
func (p Person) NameMale() string {
	return p.FirstNameMale() + " " + p.LastName()
}

// NameFemale returns random female name
func (p Person) NameFemale() string {
	return p.FirstNameFemale() + " " + p.LastName()
}

// Gender returns random gender
func (p Person) Gender() string {
	return p.Faker.RandomStringElement([]string{p.GenderMale(), p.GenderFemale()})
}

// GenderMale returns male gender
func (p Person) GenderMale() string {
	return "Male"
}

// GenderFemale returns female gender
func (p Person) GenderFemale() string {
	return "Female"
}
