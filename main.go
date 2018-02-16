package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty.usbserial", Baud: 115200, ReadTimeout: 10 * time.Millisecond}
	port, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Bifrost v0.0.1\n")
	portReader := bufio.NewReader(port)

	go func() {
		for {
			response, _ := portReader.ReadBytes('\n')
			fmt.Print(string(response))
		}
	}()

	port.Write([]byte("\r"))

	for {
		stdinReader := bufio.NewReader(os.Stdin)
		input, _ := stdinReader.ReadString('\n')
		port.Write([]byte(input))
	}
}
