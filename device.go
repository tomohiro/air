package main

import (
	"fmt"

	"github.com/Tomohiro/go-airplay"
)

// Devices Returns AirPlay devices
func Devices() []airplay.Device {
	devices := airplay.Devices()
	if len(devices) == 0 {
		fmt.Println("air: AirPlay Devices not found")
	}

	return devices
}
