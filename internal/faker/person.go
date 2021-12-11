package faker

import (
	"strings"
)

type Person struct {
	Faker *Faker
}

func (p Person) FirstName() string {
	names := make([]string, len(p.Faker.FirstNameMale)+len(p.Faker.FirstNameFemale))
	i := 0

	for j := 0; j < len(p.Faker.FirstNameMale); j++ {
		names[i] = p.Faker.FirstNameMale[j]
		i++
	}

	for j := 0; j < len(p.Faker.FirstNameFemale); j++ {
		names[i] = p.Faker.FirstNameFemale[j]
		i++
	}

	return p.Faker.RandomStringElement(names)
}

func (p Person) LastName() string {
	i := p.Faker.IntBetween(0, len(p.Faker.LastName)-1)
	return p.Faker.LastName[i]
}

func (p Person) FirstNameMale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.FirstNameMale)-1)
	return p.Faker.FirstNameMale[i]
}

func (p Person) FirstNameFemale() string {
	i := p.Faker.IntBetween(0, len(p.Faker.FirstNameFemale)-1)
	return p.Faker.FirstNameFemale[i]
}

func (p Person) Name() string {
	formats := make([]string, len(p.Faker.MaleNameFormats)+len(p.Faker.FemaleNameFormats))
	i := 0

	for j := 0; j < len(p.Faker.MaleNameFormats); j++ {
		formats[i] = p.Faker.MaleNameFormats[j]
		i++
	}

	for j := 0; j < len(p.Faker.FemaleNameFormats); j++ {
		formats[i] = p.Faker.FemaleNameFormats[j]
		i++
	}

	name := formats[p.Faker.IntBetween(0, len(formats)-1)]

	// {{FirstNameMale}}
	if strings.Contains(name, "{{FirstNameMale}}") {
		name = strings.ReplaceAll(name, "{{FirstNameMale}}", p.FirstNameMale())
	}

	// {{FirstNameFemale}}
	if strings.Contains(name, "{{FirstNameFemale}}") {
		name = strings.ReplaceAll(name, "{{FirstNameFemale}}", p.FirstNameFemale())
	}

	// {{LastName}}
	if strings.Contains(name, "{{LastName}}") {
		name = strings.ReplaceAll(name, "{{LastName}}", p.LastName())
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
