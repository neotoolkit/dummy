package faker

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
