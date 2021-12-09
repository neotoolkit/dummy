package faker

import (
	"github.com/google/uuid"
)

type UUID struct {
	Faker *Faker
}

func (u UUID) V4() string {
	return uuid.New().String()
}
