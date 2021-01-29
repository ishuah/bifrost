package main

import (
	"bufio"
	"fmt"
	"io"
	"time"

	//"github.com/tarm/serial"
	"github.com/pkg/term"
)

// Connect contains all the configuration necessary
// to open a serial port
// type Connect struct {
// 	config     *serial.Config
// 	port       *serial.Port
// 	portReader *bufio.Reader
// 	portChan   chan []byte
// 	stateChan  chan error
// }

type Connect struct {
	portPath   string
	baudRate   int
	port       *term.Term
	portReader *bufio.Reader
	stateChan  chan error
}

// NewConnection returns a pointer to a Connect instance
func NewConnection(portPath string, baudRate int) (*Connect, error) {
	//config := serial.Config{Name: portPath, Baud: baudRate, ReadTimeout: time.Nanosecond}
	// port, err := serial.OpenPort(&config)

	t := term.Speed(baudRate)
	port, err := term.Open(portPath, t)
	if err != nil {
		return nil, err
	}
	port.SetRaw()
	portReader := bufio.NewReader(port)
	stateChan := make(chan error)
	return &Connect{portPath: portPath, baudRate: baudRate, port: port,
		portReader: portReader,
		stateChan:  stateChan}, nil
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
		port, err := term.Open(c.portPath)
		if err != nil {
			continue
		}
		c.port = port
		c.portReader = bufio.NewReader(port)
		c.stateChan <- nil
		return
	}
}

func (c *Connect) read() {
	for {
		response, err := c.portReader.ReadBytes('\n')
		// report the error
		if err != nil && err != io.EOF {
			c.stateChan <- err
			return
		}
		if len(response) > 0 {
			fmt.Print(string(response))
		}
	}
}

func (c *Connect) Write(message []byte) {
	_, err := c.port.Write(message)
	if err != nil {
		fmt.Printf("Error writing to serial port: %v ", err)
	}
}
