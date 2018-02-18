package screen

import (
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer           []string
	width, height    int
	x, y             int
	cursorX, cursorY int
	input            []rune
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
	termbox.SetCell(s.x, s.y, ' ', termbox.ColorWhite, termbox.ColorWhite)
}

func (s *Screen) Write(line string) {
	defer termbox.Flush()
	s.buffer = append(s.buffer, line)
	lines := []string{line}
	if s.y > s.height {
		s.y = 0
		s.x = 0
		scope := (len(s.buffer) - s.height) + 1
		lines = s.buffer[scope:]
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, line := range lines {
		s.write(line)
	}
}

func (s *Screen) updateInput() {
	inputX := s.x
	inputY := s.y
	defer termbox.Flush()
	for x := 0; x <= len(s.input)+1; x++ {
		termbox.SetCell(x+inputX, inputY, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, c := range s.input {
		termbox.SetCell(inputX, inputY, c, termbox.ColorDefault, termbox.ColorDefault)
		inputX++
	}
	s.cursorX = inputX
	s.cursorY = inputY
	s.drawCursor()
}

func (s *Screen) drawCursor() {
	termbox.SetCell(s.cursorX, s.cursorY, ' ', termbox.ColorBlack, termbox.ColorWhite)
	termbox.Flush()
}

func (s *Screen) clearInput() {
	s.input = s.input[:0]
}

func (s *Screen) AddInputChar(char rune) {
	s.input = append(s.input, char)
	s.updateInput()
}

func (s *Screen) DeleteInputChar() {
	if len(s.input) == 0 {
		return
	}
	s.input = s.input[:len(s.input)-1]
	s.updateInput()
}

func (s *Screen) GetInput() []rune {
	defer s.clearInput()
	s.input = append(s.input, '\r')
	return s.input
}
