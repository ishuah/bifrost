//go:build !windows

package main

import (
	"github.com/pkg/term"
)

// termPort wraps term.Term to implement SerialPort interface
type termPort struct {
	*term.Term
}

// NewConnection returns a pointer to a Connect instance
func NewConnection(portPath string, baudRate int) (*Connect, error) {
	t := term.Speed(baudRate)
	port, err := term.Open(portPath, t)
	if err != nil {
		return nil, err
	}
	port.SetRaw()

	openPort := func() (SerialPort, error) {
		p, err := term.Open(portPath, t)
		if err != nil {
			return nil, err
		}
		p.SetRaw()
		return &termPort{p}, nil
	}

	stateChan := make(chan error)
	return &Connect{
		portPath:  portPath,
		baudRate:  baudRate,
		port:      &termPort{port},
		stateChan: stateChan,
		openPort:  openPort,
	}, nil
}
