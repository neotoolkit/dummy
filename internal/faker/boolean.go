package faker

type Boolean struct {
	Faker *Faker
}

func (b Boolean) Boolean() bool {
	return b.Faker.IntBetween(0, 100) > 50
}
