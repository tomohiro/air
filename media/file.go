package media

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// File is local file
type File struct {
	Path string
}

// FileServerAddr is ip and port of serving files
var FileServerAddr string

func init() {
	ip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}

	port, err := findFreePort()
	if err != nil {
		log.Fatal(err)
	}

	FileServerAddr = fmt.Sprintf("%s:%s", ip, port)
}

// NewFile creates a new file
func NewFile(path string) *File {
	file := new(File)

	http.HandleFunc(fmt.Sprintf("/%s", path), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, path)
	})

	file.Path = fmt.Sprintf("http://%s/%s", FileServerAddr, path)
	return file
}

// URL returns serve media url
func (m *File) URL() string {
	serve()
	return m.Path
}

func serve() {
	go http.ListenAndServe(FileServerAddr, nil)
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
