package faker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Faker is struct for Faker
type Faker struct {
	Generator        *rand.Rand
	firstNameMale    []string
	firstNameFemale  []string
	lastName         []string
	maleNameFormat   []string
	femaleNameFormat []string
	usernameFormat   []string
	// Generic top-level domain
	gTLD []string
}

// NewFaker returns a new instance of Faker instance with a random seed
func NewFaker() Faker {
	source := rand.NewSource(time.Now().Unix())

	return Faker{
		Generator: rand.New(source),
		firstNameMale: []string{
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
		firstNameFemale: []string{
			"Alice",
			"Grace",
			"Jennifer",
			"Luna",
			"Mary", "Maya", "Mia",
			"Patricia",
			"Ruby",
			"Willow",
		},
		lastName: []string{
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
		maleNameFormat: []string{
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{lastName}} {{firstNameMale}}",
		},
		femaleNameFormat: []string{
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{lastName}} {{firstNameFemale}}",
		},
		usernameFormat: []string{
			"{{lastName}}.{{firstName}}",
			"{{firstName}}.{{lastName}}",
			"{{firstName}}",
			"{{lastName}}",
		},
		gTLD: []string{
			"com",
			"info",
			"net",
			"org",
		},
	}
}

// Internet returns Internet instance
func (f Faker) Internet() Internet {
	return Internet{&f}
}

// Person returns Person instance
func (f Faker) Person() Person {
	return Person{&f}
}

// Boolean returns Boolean instance
func (f Faker) Boolean() Boolean {
	return Boolean{&f}
}

// UUID returns UUID instance
func (f Faker) UUID() UUID {
	return UUID{&f}
}

// IntBetween returns a Int between a given minimum and maximum values
func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return f.Generator.Intn(diff+1) + min
}

// RandomStringElement returns a random string element from a given list of strings
func (f Faker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}

// Asciify returns string that replace all "*" characters with random ASCII values from a given string
func (f Faker) Asciify(in string) string {
	var out strings.Builder

	for i := 0; i < len(in); i++ {
		if in[i] == '*' {
			out.WriteString(fmt.Sprintf("%c", f.IntBetween(97, 122)))
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}
