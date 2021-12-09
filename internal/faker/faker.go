package faker

import (
	"math/rand"
	"time"
)

type Faker struct {
	Generator *rand.Rand
}

func NewFaker() Faker {
	source := rand.NewSource(time.Now().Unix())

	return Faker{
		Generator: rand.New(source),
	}
}

func (f Faker) UUID() UUID {
	return UUID{&f}
}

func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return f.Generator.Intn(diff+1) + min
}
