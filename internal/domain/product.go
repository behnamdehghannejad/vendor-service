package domain

import (
	"errors"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Active      bool
	CategoryID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func IsActiveProduct(isActive bool) error {
	if !isActive {
		return errors.New("product is inactive")
	}
	return nil
}
