package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pdxgo/whispering-gophers/util"
)

// Open launches server from the media file.
func Open(name string) (string, error) {
	path, err := filepath.Abs(name)
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
		return serve(path), nil
	}

	return "", nil
}

func serve(name string) string {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set `no-cache` to flush the cache on Apple TV.
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, name)
	})

	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(l, http.DefaultServeMux)
	return fmt.Sprintf("http://%s", l.Addr().String())
}
