package main

import (
	"fmt"
	"io"
	"time"
)

// SerialPort defines the interface for platform-specific serial port implementations
type SerialPort interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
}

// Connect manages a serial port connection with auto-reconnect capability
type Connect struct {
	portPath  string
	baudRate  int
	port      SerialPort
	stateChan chan error
	openPort  func() (SerialPort, error)
}

// Start initializes a read loop that attempts to reconnect
// when the connection is broken
func (c *Connect) Start() {
	go c.read()
	for {
		select {
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
	for {
		time.Sleep(time.Second)
		port, err := c.openPort()
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
		if err != nil && err != io.EOF {
			c.stateChan <- err
			return
		}
		if n > 0 {
			fmt.Print(string(buf))
		}
	}
}

// Write sends data to the serial port
func (c *Connect) Write(message []byte) {
	_, err := c.port.Write(message)
	if err != nil {
		fmt.Printf("Error writing to serial port: %v ", err)
	}
}
