package domain

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
