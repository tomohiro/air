package main

import (
	"errors"
	"os"
	"path/filepath"
)

type media interface {
	URL() string
}

var (
	isFile     = "file"
	isResource = "resource"
)

func classifyType(path string) (string, error) {
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
		return isFile, nil
	}

	return "", nil
}
