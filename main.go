package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

const version = "v0.1.21-rc1"

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

	connect, err := NewConnection(portPath, baud)
	if err != nil {
		log.Printf("FatalError: %v", err)
		return
	}

	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	screen := NewScreen()
	// Welcome message
	screen.Write(welcomeMessage(portPath, baud))

	screenChan := make(chan []byte)
	go connect.Start(screenChan)
	go screen.BufferedWriter(screenChan)

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Ch != 0 || ev.Key == termbox.KeySpace {
				char := ev.Ch
				if ev.Key == termbox.KeySpace {
					char = ' '
				}
				connect.Write([]byte(string(char)))
			} else {
				switch ev.Key {
				case termbox.KeyEsc:
					connect.Write([]byte{'\x1b'})
				case termbox.KeyCtrlBackslash:
					return
				case termbox.KeyTab:
					connect.Write([]byte{'\x09'})
				case termbox.KeyCtrlC:
					connect.Write([]byte{'\x03'})
				case termbox.KeyEnter:
					connect.Write([]byte{'\r'})
				case termbox.KeyBackspace:
					connect.Write([]byte{'\x7F'})
				case termbox.KeyBackspace2:
					connect.Write([]byte{'\x7F'})
				case termbox.KeyArrowLeft:
					connect.Write([]byte{'\x1b', '[', 'D'})
				case termbox.KeyArrowRight:
					connect.Write([]byte{'\x1b', '[', 'C'})
				case termbox.KeyArrowUp:
					connect.Write([]byte{'\x1b', '[', 'A'})
				case termbox.KeyArrowDown:
					connect.Write([]byte{'\x1b', '[', 'B'})
				}
			}
		}
	}
}
