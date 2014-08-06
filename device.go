package main

import (
	"errors"

	"github.com/Tomohiro/go-airplay"
)

// Devices returns AirPlay devices
func Devices() ([]airplay.Device, error) {
	devices := airplay.Devices()
	if len(devices) == 0 {
		return devices, errors.New("air: AirPlay Devices not found")
	}

	return devices, nil
}
