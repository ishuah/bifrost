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

const version = "v0.1.20-rc1"

var header = fmt.Sprintf("\nBifrost %s\n", version)
var helpText = fmt.Sprintf(`%s
Bifrost is a tiny terminal emulator for serial port communication.

    Usage:
	  bifrost [flags]
	  
    Flags:
      -port-path	Name/path of the serial port
      -baud		The baud rate to use on the connection
      -help		This help message
	`, header)

func welcomeMessage(portPath string, baud int) string {
	return fmt.Sprintf(`%s
Options:
    Port:		%s
    Baud rate:	%d

Press Ctrl+\ to exit
		`, header, portPath, baud)
}

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
	var portPath string
	var baud int
	var help bool
	flag.StringVar(&portPath, "port-path", "/dev/tty.usbserial", "Name/path of the serial port")
	flag.IntVar(&baud, "baud", 115200, "The baud rate to use on the connection")
	flag.BoolVar(&help, "help", false, "A brief help message")
	flag.Parse()

	if help {
		fmt.Println(helpText)
		return
	}

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	screen := screen.NewScreen()

	c := &serial.Config{Name: portPath, Baud: baud, ReadTimeout: time.Nanosecond}
	port, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	portReader := bufio.NewReader(port)

	// Welcome message
	screen.Write(welcomeMessage(portPath, baud))

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
