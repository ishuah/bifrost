package main

import (
	"bufio"
	"fmt"
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
	config := &serial.Config{Name: portPath, Baud: baudRate, ReadTimeout: time.Nanosecond}
	port, err := serial.OpenPort(config)
	if err != nil {
		return &Connect{}, err
	}
	portReader := bufio.NewReader(port)
	stateChan := make(chan error)
	return &Connect{config: config, port: port,
		portReader: portReader,
		stateChan:  stateChan}, nil
}

func (c *Connect) Start(screenChan chan []byte) {
	go c.read(screenChan)
	for {
		select {
		case err := <-c.stateChan:
			if err != nil {
				screenChan <- []byte(fmt.Sprintf("\nError connecting to %s", c.config.Name))
				go c.initialize()
			} else {
				screenChan <- []byte(fmt.Sprintf("\nConnection to %s reestablished!", c.config.Name))
				go c.read(screenChan)
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

func (c *Connect) read(screenChan chan []byte) {
	for {
		response, err := c.portReader.ReadBytes('\n')
		// report the error
		if err != nil && err.Error() != "EOF" {
			c.stateChan <- err
			return
		}
		if len(response) > 0 {
			screenChan <- response
		}
	}
}

func (c *Connect) Write(message []byte) {
	c.port.Write(message)
}
