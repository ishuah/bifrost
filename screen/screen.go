package screen

import (
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer                                []string
	width, height, x, y, promptX, promptY int
	input                                 []rune
}

func NewScreen() Screen {
	width, height := termbox.Size()
	return Screen{width: width, height: height}
}

func (s *Screen) write(line string) {
	defer termbox.Flush()
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
	s.buffer = append(s.buffer, line)
	lines := []string{line}
	if s.y > s.height {
		s.y = 0
		s.x = 0
		lines = s.buffer[len(s.buffer)-s.height:]
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, line := range lines {
		s.write(line)
	}
}

func (s *Screen) updatePrompt() {
	s.promptX = s.x
	s.promptY = s.y
	defer termbox.Flush()
	for x := 0; x <= len(s.input)+1; x++ {
		termbox.SetCell(x+s.promptX, s.promptY, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, c := range s.input {
		termbox.SetCell(s.promptX, s.promptY, c, termbox.ColorDefault, termbox.ColorDefault)
		s.promptX++
	}
}

func (s *Screen) clearPrompt() {
	s.input = s.input[:0]
}

func (s *Screen) AddPromptChar(char rune) {
	s.input = append(s.input, char)
	s.updatePrompt()
}

func (s *Screen) DeletePromptChar() {
	s.input = s.input[:len(s.input)-1]
	s.updatePrompt()
}

func (s *Screen) PromptInput() []rune {
	defer s.clearPrompt()
	s.input = append(s.input, '\r')
	return s.input
}
