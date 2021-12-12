package faker

import (
	"github.com/google/uuid"
)

// UUID is struct for UUID
type UUID struct {
	Faker *Faker
}

// V4 returns UUID V4 as string
func (u UUID) V4() string {
	return uuid.New().String()
}
