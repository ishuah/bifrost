//go:build windows

package main

import (
	"fmt"
	"io"
	"time"

	"github.com/tarm/serial"
)

type Connect struct {
	portPath string
	baudRate int
	port	 *serial.Port
	stateChan chan error
}

// NewConnection returns a pointer to a Connect instance
func NewConnection(portPath string, baudRate int) (*Connect, error) {
	c := &serial.Config{Name: portPath, Baud: baudRate}
	port, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	stateChan := make(chan error)
	return &Connect{portPath: portPath, baudRate: baudRate, port: port,
		stateChan: stateChan}, nil
}

// Start initializes a read loop that attempts to reconnect
// when the connection is broken
func (c *Connect) Start() {
	go c.read()
	for {
		select{
		case err := <-c.stateChan:
			if err != nil {
				fmt.Printf("Error connecting to %s", c.portPath)
				go c.initialize()
			} else {
				fmt.Printf(" | Connection to %s reestablished!", c.portPath)
				go c.read()
			}
		}
	}
}

func (c *Connect) initialize() {
	c.port.Close()
	config := &serial.Config{Name: c.portPath, Baud: c.baudRate}
	for {
		time.Sleep(time.Second)
		port, err := serial.OpenPort(config)
		if err != nil {
			continue
		}
		c.port = port
		c.stateChan <- nil
		return
	}
}

func (c *Connect) read() {
	buf := make([]byte, 1)
	for {
		n, err := io.ReadFull(c.port, buf)
		// report the error
		if err != nil && err != io.EOF {
			c.stateChan <- err
			return
		}
		if n > 0 {
			fmt.Print(string(buf))
		}
	}
}

func (c *Connect) Write(message []byte) {
	_, err := c.port.Write(message)
	if err != nil {
		fmt.Printf("Error writing to serial port: %v ", err)
	}
}