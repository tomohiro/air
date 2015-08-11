package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type file struct {
	Path string
}

func newFile(path string) *file {
	return &file{Path: path}
}

// URL returns serve media url
func (m *file) URL() string {
	return serve(m.Path)
}

func serve(file string) string {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, file)
	})

	ip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}
	port, err := findFreePort()
	if err != nil {
		return ""
	}
	go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return fmt.Sprintf("http://%s:%s", ip, port)
}

// https://code.google.com/p/whispering-gophers/source/browse/util/helper.go
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func findFreePort() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	port := strings.Split(l.Addr().String(), ":")[1]
	return port, nil
}
