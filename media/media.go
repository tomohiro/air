package media

import (
	"fmt"
	"os"
	"path/filepath"
)

// Media file interface
type Media interface {
	URL() string
}

var (
	// IsDirectory is media type is directory
	IsDirectory = "directory"

	// IsFile is media type is file
	IsFile = "file"

	// IsResource is media type is resource
	IsResource = "resource"
)

// IsSupported check if the given path is supported file type.
// If the given path unsupported file type, it's returns error.
func IsSupported(path string) error {
	return fmt.Errorf("%s is unsupported mime type", path)
}

// Classify classify the given path to media type.
// Types are file, directory, resource.
func Classify(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	stat, err := f.Stat()
	if err != nil {
		return "", err
	}

	switch mode := stat.Mode(); {
	case mode.IsDir():
		return IsDirectory, nil
	case mode.IsRegular():
		return IsFile, nil
	}

	return "", nil
}
