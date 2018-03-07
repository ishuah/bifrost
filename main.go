package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ishuah/bifrost/screen"
	"github.com/nsf/termbox-go"
	"github.com/tarm/serial"
)

const version = "v0.1.20"

func bufferedReader(portReader *bufio.Reader, buf chan []byte) {
	for {
		response, _ := portReader.ReadBytes('\n')
		if len(response) > 0 {
			buf <- response
		}
	}
}

func bufferedWriter(screen screen.Screen, buf chan []byte) {
	for {
		select {
		case response := <-buf:
			screen.Write(string(response))
		}
	}
}

func main() {
	name := flag.String("name", "/dev/tty.usbserial", "serial device port")
	baud := flag.Int("baud", 115200, "baud rate")

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	screen := screen.NewScreen()

	c := &serial.Config{Name: *name, Baud: *baud, ReadTimeout: time.Nanosecond}
	port, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	portReader := bufio.NewReader(port)

	// Welcome message
	screen.Write(fmt.Sprintf("\nBifrost %s\n\n", version))
	screen.Write("Options:\n")
	screen.Write(fmt.Sprintf("\t\tPort: %s\n\t\tBaud rate: %d\n\n", *name, *baud))
	screen.Write("Press Ctrl+\\ to exit\n\n")

	buf := make(chan []byte)
	go bufferedReader(portReader, buf)
	go bufferedWriter(screen, buf)

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Ch != 0 || ev.Key == termbox.KeySpace {
				char := ev.Ch
				if ev.Key == termbox.KeySpace {
					char = ' '
				}
				port.Write([]byte(string(char)))
			} else {
				switch ev.Key {
				case termbox.KeyEsc:
					port.Write([]byte{'\x1b'})
				case termbox.KeyCtrlBackslash:
					return
				case termbox.KeyTab:
					port.Write([]byte{'\x09'})
				case termbox.KeyCtrlC:
					port.Write([]byte{'\x03'})
				case termbox.KeyEnter:
					port.Write([]byte{'\r'})
				case termbox.KeyBackspace:
					port.Write([]byte{'\x7F'})
				case termbox.KeyBackspace2:
					port.Write([]byte{'\x7F'})
				case termbox.KeyArrowLeft:
					port.Write([]byte{'\x1b', '[', 'D'})
				case termbox.KeyArrowRight:
					port.Write([]byte{'\x1b', '[', 'C'})
				case termbox.KeyArrowUp:
					port.Write([]byte{'\x1b', '[', 'A'})
				case termbox.KeyArrowDown:
					port.Write([]byte{'\x1b', '[', 'B'})
				}
			}
		}
	}
}
