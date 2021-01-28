package main

import (
	"flag"
	"fmt"
	"log"
)

const version = "v0.2.1"

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
    Port:	%s
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
	// Welcome message
	fmt.Print(welcomeMessage(portPath, baud))

	go connect.Start()

	for {
		key := pollKeyEvents()

		if len(key.Value) != 0 {
			connect.Write(key.Value)
		} else {
			switch key.Type {
			case Esc:
				connect.Write([]byte{'\x1b'})
			case CtrlBackslash:
				fmt.Print("bye!")
				return
			case Tab:
				connect.Write([]byte{'\x09'})
			case CtrlC:
				connect.Write([]byte{'\x03'})
			case Enter:
				connect.Write([]byte{'\r'})
			case Backspace:
				connect.Write([]byte{'\x7F'})
			case Delete:
				connect.Write([]byte{'\x1b', '[', '3', '~'})
			case LeftArrow:
				connect.Write([]byte{'\x1b', '[', 'D'})
			case RightArrow:
				connect.Write([]byte{'\x1b', '[', 'C'})
			case UpArrow:
				connect.Write([]byte{'\x1b', '[', 'A'})
			case DownArrow:
				connect.Write([]byte{'\x1b', '[', 'B'})
			case Space:
				connect.Write([]byte{' '})
			}
		}
	}
}
