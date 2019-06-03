package main

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/tarm/serial"
)

type Connect struct {
	config     *serial.Config
	port       *serial.Port
	portReader *bufio.Reader
	portChan   chan []byte
	stateChan  chan error
}

func NewConnection(portPath string, baudRate int) (*Connect, error) {
	config := serial.Config{Name: portPath, Baud: baudRate, ReadTimeout: time.Nanosecond}
	port, err := serial.OpenPort(&config)
	if err != nil {
		return nil, err
	}
	portReader := bufio.NewReader(port)
	stateChan := make(chan error)
	return &Connect{config: &config, port: port,
		portReader: portReader,
		stateChan:  stateChan}, nil
}

func (c *Connect) Start() {
	go c.read()
	for {
		select {
		case err := <-c.stateChan:
			if err != nil {
				fmt.Printf("Error connecting to %s", c.config.Name)
				go c.initialize()
			} else {
				fmt.Printf(" | Connection to %s reestablished!", c.config.Name)
				go c.read()
			}
		}
	}
}

func (c *Connect) initialize() {
	c.port.Close()
	for {
		time.Sleep(time.Second)
		port, err := serial.OpenPort(c.config)
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
