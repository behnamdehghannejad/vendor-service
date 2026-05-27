package domain

import (
	"errors"
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

func IsActiveVendor(isActive bool) error {
	if !isActive {
		return errors.New("vendor is inactive")
	}
	return nil
}
