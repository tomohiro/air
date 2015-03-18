package media

import (
	"errors"
	"os"
	"path/filepath"
)

// Media file interface
type Media interface {
	URL() string
}

var (
	// IsFile is media type is file
	IsFile = "file"

	// IsResource is media type is resource
	IsResource = "resource"
)

// ClassifyType classify type from path
func ClassifyType(path string) (string, error) {
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
		return "", errors.New("directory is unsupported")
	case mode.IsRegular():
		return IsFile, nil
	}

	return "", nil
}
