package domain

import (
	"errors"
	"time"
)

type Product struct {
	ID                 int
	DiscountPercentage float64
	Name               string
	Description        string
	Active             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func IsActiveProduct(isActive bool) error {
	if !isActive {
		return errors.New("vendor is inactive")
	}
	return nil
}
