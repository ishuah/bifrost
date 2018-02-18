package main

import (
	"bufio"
	"log"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/tarm/serial"
)

type Screen struct {
	buffer              []string
	width, height, x, y int
}

func (s *Screen) Write(content string) {
	defer termbox.Flush()
	s.buffer = append(s.buffer, content)
	lines := []string{content}
	if s.y > s.height {
		s.y = 0
		s.x = 0
		lines = s.buffer[len(s.buffer)-s.height:]
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, line := range lines {
		for _, char := range line {
			if s.x > s.width {
				s.x = 0
				s.y++
			}
			// new line character
			if char == 10 {
				s.x = 0
				s.y++
				continue
			}
			termbox.SetCell(s.x, s.y, char, termbox.ColorDefault, termbox.ColorDefault)
			s.x++
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()
	width, height := termbox.Size()
	screen := Screen{width: width, height: height, y: 0}

	c := &serial.Config{Name: "/dev/tty.usbserial", Baud: 115200, ReadTimeout: 10 * time.Millisecond}
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
				screen.Write(string(char))
			} else {
				switch ev.Key {
				case termbox.KeyEsc:
					return
				}
			}
		}
	}
}
