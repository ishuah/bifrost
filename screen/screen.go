package screen

import (
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer              []string
	width, height, x, y int
	input               []rune
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

func (s *Screen) updatePrompt() {
	promptX := s.x
	promptY := s.y
	defer termbox.Flush()
	for x := 0; x <= len(s.input)+1; x++ {
		termbox.SetCell(x+promptX, promptY, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, c := range s.input {
		termbox.SetCell(promptX, promptY, c, termbox.ColorDefault, termbox.ColorDefault)
		promptX++
	}
	termbox.SetCell(promptX, promptY, ' ', termbox.ColorWhite, termbox.ColorWhite)
}

func (s *Screen) clearPrompt() {
	s.input = s.input[:0]
}

func (s *Screen) AddPromptChar(char rune) {
	s.input = append(s.input, char)
	s.updatePrompt()
}

func (s *Screen) DeletePromptChar() {
	if len(s.input) == 0 {
		return
	}
	s.input = s.input[:len(s.input)-1]
	s.updatePrompt()
}

func (s *Screen) PromptInput() []rune {
	defer s.clearPrompt()
	s.input = append(s.input, '\r')
	return s.input
}
