package utils

import (
	"os"
	"path/filepath"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
)

const rootPathKey = "ROOT_PATH"

func SetRootPath(path string) error {
	root, err := filepath.Abs(path)
	if err != nil {
		return apperror.Wrap(err).
			Warning().
			Forbidden().
			Build()
	}
	os.Setenv(rootPathKey, root)
	return nil
}

func GetRootPath() string {
	return os.Getenv(rootPathKey)
}
