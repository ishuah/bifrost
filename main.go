package main

import (
	"bufio"
	"log"
	"time"

	"github.com/ishuah/bifrost/screen"
	"github.com/nsf/termbox-go"
	"github.com/tarm/serial"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	screen := screen.NewScreen()

	c := &serial.Config{Name: "/dev/tty.usbserial", Baud: 115200, ReadTimeout: 500 * time.Nanosecond}
	port, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	screen.Write("Bifrost alpha build\n")
	port.Write([]byte("\r"))

	portReader := bufio.NewReader(port)

	go func() {
		for {
			response, _ := portReader.ReadBytes('\n')
			if len(response) > 0 {
				screen.Write(string(response))
			}
		}
	}()

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Ch != 0 || ev.Key == termbox.KeySpace {
				char := ev.Ch
				if ev.Key == termbox.KeySpace {
					char = ' '
				}
				screen.AddPromptChar(char)
			} else {
				switch ev.Key {
				case termbox.KeyEsc:
					return
				case termbox.KeyEnter:
					port.Write([]byte(string(screen.PromptInput())))
				case termbox.KeyBackspace2:
					screen.DeletePromptChar()
				}
			}
		}
	}
}
