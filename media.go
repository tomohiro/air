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

// Open launches server from the media file path.
func Open(path string) (string, error) {
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
		return serve(path), nil
	}

	return "", nil
}

// serve listens the file f from the local file system as local network address
func serve(f string) string {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, f)
	})

	laddr := availableAddr()
	go http.ListenAndServe(laddr, nil)
	return fmt.Sprintf("http://%s", laddr)
}

// availableAddr returns an address that available non-lookback IPv4 network
// interface and port.
func availableAddr() string {
	ln, err := util.Listen()
	defer ln.Close()
	if err != nil {
		log.Fatal(err)
	}
	return ln.Addr().String()
}
