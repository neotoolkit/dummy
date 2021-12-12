package faker

import "strings"

// Internet is struct for Internet
type Internet struct {
	Faker *Faker
}

// Username returns random username
func (i Internet) Username() string {
	username := i.Faker.RandomStringElement(i.Faker.usernameFormat)

	p := i.Faker.Person()

	// {{firstName}}
	if strings.Contains(username, "{{firstName}}") {
		username = strings.ReplaceAll(username, "{{firstName}}", strings.ToLower(p.FirstName()))
	}

	// {{lastName}}
	if strings.Contains(username, "{{lastName}}") {
		username = strings.ReplaceAll(username, "{{lastName}}", strings.ToLower(p.LastName()))
	}

	return username
}

// GTLD returns random generic top-level domain
func (i Internet) GTLD() string {
	return i.Faker.RandomStringElement(i.Faker.gTLD)
}

// Domain returns random domain
func (i Internet) Domain() string {
	return strings.Join([]string{i.Faker.Asciify("***"), i.GTLD()}, ".")
}

// Email returns random email
func (i Internet) Email() string {
	return i.Username() + "@" + i.Domain()
}
