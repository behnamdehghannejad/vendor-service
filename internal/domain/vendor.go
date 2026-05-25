package domain

import (
	"time"
)

type Vendor struct {
	ID        int
	Code      string
	Name      string
	Email     string
	Phone     string
	Address   string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
