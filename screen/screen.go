package screen

import (
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer        []string
	width, height int
	x, y          int
}

func NewScreen() Screen {
	width, height := termbox.Size()
	return Screen{width: width, height: height}
}

func (s *Screen) write(line string) {
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

func (s *Screen) Write(line string) {
	defer termbox.Flush()
	s.buffer = append(s.buffer, line)
	lines := []string{line}

	if s.y > s.height {
		s.x = 0
		s.y = 0
		scope := (len(s.buffer) - s.height) + 1
		lines = s.buffer[scope:]
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, line := range lines {
		s.write(line)
	}
	termbox.SetCell(s.x, s.y, ' ', termbox.ColorBlack, termbox.ColorWhite)
}
