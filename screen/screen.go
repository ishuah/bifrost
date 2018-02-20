package screen

import (
	"github.com/ishuah/bifrost/command"
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer        []string
	width, height int
	prompt        *command.Prompt
}

func NewScreen() Screen {
	width, height := termbox.Size()
	prompt := command.NewPrompt(width)
	return Screen{width: width, height: height, prompt: &prompt}
}

func (s *Screen) Write(line string) {
	defer termbox.Flush()
	s.buffer = append(s.buffer, line)
	lines := []string{line}
	_, y := s.prompt.GetPosition()
	if y > s.height {
		s.prompt.Reset()
		scope := (len(s.buffer) - s.height) + 1
		lines = s.buffer[scope:]
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	for _, line := range lines {
		s.prompt.Write(line)
	}
	s.prompt.UpdatePosition()
	s.prompt.Draw()
}

func (s *Screen) InsertInputChar(char rune) {
	s.prompt.InsertInputChar(char)
}

func (s *Screen) DeleteInputChar() {
	s.prompt.DeleteInputChar()
}

func (s *Screen) ReturnInput() []rune {
	return s.prompt.ReturnInput()
}
