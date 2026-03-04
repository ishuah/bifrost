//go:build windows

package main

import (
	"github.com/tarm/serial"
)

// NewConnection returns a pointer to a Connect instance
func NewConnection(portPath string, baudRate int) (*Connect, error) {
	config := &serial.Config{Name: portPath, Baud: baudRate}
	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	openPort := func() (SerialPort, error) {
		return serial.OpenPort(config)
	}

	stateChan := make(chan error)
	return &Connect{
		portPath:  portPath,
		baudRate:  baudRate,
		port:      port,
		stateChan: stateChan,
		openPort:  openPort,
	}, nil
}
