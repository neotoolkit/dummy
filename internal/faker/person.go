package faker

import (
	"strings"
)

type Person struct {
	Faker *Faker
}

func (p Person) FirstName() string {
	firstNames := make([]string, len(p.Faker.firstNameMale)+len(p.Faker.firstNameFemale))
	i := 0

	for j := 0; j < len(p.Faker.firstNameMale); j++ {
		firstNames[i] = p.Faker.firstNameMale[j]
		i++
	}

	for j := 0; j < len(p.Faker.firstNameFemale); j++ {
		firstNames[i] = p.Faker.firstNameFemale[j]
		i++
	}

	return p.Faker.RandomStringElement(firstNames)
}

func (p Person) LastName() string {
	i := p.Faker.IntBetween(0, len(p.Faker.lastName)-1)
	return p.Faker.lastName[i]
}

func (p Person) FirstNameMale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.firstNameMale)-1)
	return p.Faker.firstNameMale[i]
}

func (p Person) FirstNameFemale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.firstNameFemale)-1)
	return p.Faker.firstNameFemale[i]
}

func (p Person) Name() string {
	formats := make([]string, len(p.Faker.maleNameFormat)+len(p.Faker.femaleNameFormat))
	i := 0

	for j := 0; j < len(p.Faker.maleNameFormat); j++ {
		formats[i] = p.Faker.maleNameFormat[j]
		i++
	}

	for j := 0; j < len(p.Faker.femaleNameFormat); j++ {
		formats[i] = p.Faker.femaleNameFormat[j]
		i++
	}

	name := formats[p.Faker.IntBetween(0, len(formats)-1)]

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

func (p Person) NameMale() string {
	return p.FirstNameMale() + " " + p.LastName()
}

func (p Person) NameFemale() string {
	return p.FirstNameFemale() + " " + p.LastName()
}

func (p Person) Gender() string {
	return p.Faker.RandomStringElement([]string{p.GenderMale(), p.GenderFemale()})
}

func (p Person) GenderMale() string {
	return "Male"
}

func (p Person) GenderFemale() string {
	return "Female"
}
