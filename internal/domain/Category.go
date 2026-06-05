package domain

import "errors"

type Category struct {
	ID       int
	Name     string
	Active   bool
	ParentID *int
	Path     string
}

func IsActiveCategory(isActive bool) error {
	if !isActive {
		return errors.New("category is inactive")
	}
	return nil
}

func ValidatePathFormat(path string) error {
	return nil // todo fix this later
}
