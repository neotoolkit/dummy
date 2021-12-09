package faker

import (
	"math/rand"
	"time"
)

type Faker struct {
	Generator       *rand.Rand
	FirstNameMale   []string
	FirstNameFemale []string
	LastName        []string
}

func NewFaker() Faker {
	source := rand.NewSource(time.Now().Unix())

	return Faker{
		Generator: rand.New(source),
		FirstNameMale: []string{
			"Alexander", "Anthony",
			"Daniel",
			"Elon",
			"Freddie",
			"James", "John",
			"Leo",
			"Matthew",
			"Oliver",
			"Robert",
			"Sergey",
		},
		FirstNameFemale: []string{
			"Alice",
			"Grace",
			"Jennifer",
			"Luna",
			"Mary", "Maya", "Mia",
			"Patricia",
			"Ruby",
			"Willow",
		},
		LastName: []string{
			"Adams", "Anderson",
			"Brin", "Brown",
			"Carter", "Clarke",
			"Evans",
			"Fisher", "Fletcher", "Ford", "Fox",
			"Green",
			"Jackson", "Johnson", "Jones",
			"Lewis",
			"Miller", "Musk",
			"Owen",
			"Pike",
			"Smith",
			"Walker", "Williams",
		},
	}
}

func (f Faker) Person() Person {
	return Person{&f}
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

func (f Faker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}
