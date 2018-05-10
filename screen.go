package main

import (
	"github.com/ishuah/ansi"
	termbox "github.com/nsf/termbox-go"
)

type Screen struct {
	buffer        []string
	width, height int
	x, y          int
	fg            termbox.Attribute
}

func NewScreen() Screen {
	width, height := termbox.Size()
	width--
	height--
	return Screen{width: width, height: height, fg: termbox.ColorDefault}
}

func (s *Screen) write(line string) {
	var lex ansi.Lexer
	lex.Init([]byte(line))

	for item := lex.NextItem(); item.T != ansi.EOF; item = lex.NextItem() {
		switch item.T {
		case ansi.ControlSequence:
			seq, _ := ansi.ParseControlSequence(item.Value)
			switch seq.Prefix {
			case ansi.ControlSequenceIntroducer:
				switch seq.Command {
				case ansi.SelectGraphicRendition:
					s.graphicHandler(seq.Params)
				case ansi.CursorPosition, ansi.CursorUp,
					ansi.CursorDown, ansi.CursorForward, ansi.CursorBack:
					s.cursorHandler(seq.Command, seq.Params)
				case ansi.EraseInDisplay, ansi.EraseInLine:
					s.eraseHandler(seq.Command, seq.Params)
				}
			}
		case ansi.RawBytes:
			for _, char := range item.Value {
				if s.x > s.width {
					s.x = 0
					s.y++
				}

				switch char {
				case 7:
					continue
				case 8:
					s.x--
				case 10:
					s.x = 0
					s.y++
				case 13:
					s.x = 0
				default:
					termbox.SetCell(s.x, s.y, rune(char), s.fg, termbox.ColorDefault)
					s.x++
				}
			}
		}
	}
}

func (s *Screen) writeLines(lines []string) {
	for _, line := range lines {
		s.write(line)
	}
	termbox.SetCursor(s.x, s.y)
}

func (s *Screen) eraseHandler(command byte, params []int) {
	switch command {
	case ansi.EraseInDisplay:
		// Erase characters starting from the cursor to the end of screen
		for i := s.y; i <= s.height; i++ {
			// start from column 0 unless i is cursor row, then start from s.x so as
			// not to delete the cursor prompt and/or input
			j := 0
			if i == s.y {
				j = s.x
			}
			for j <= s.width {
				termbox.SetCell(j, i, ' ', termbox.ColorDefault, termbox.ColorDefault)
				j++
			}
		}
	case ansi.EraseInLine:
		for j := s.x; j <= s.width; j++ {
			termbox.SetCell(j, s.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (s *Screen) graphicHandler(params []int) {
	switch params[0] {
	case 0:
		s.fg = termbox.ColorDefault
	case 1:
		s.fg = termbox.AttrBold
	case 4:
		s.fg = termbox.AttrUnderline
	case 7:
		s.fg = termbox.AttrReverse
	}

	if len(params) == 1 {
		return
	}

	switch params[1] {
	case 30:
		s.fg |= termbox.ColorBlack
	case 31:
		s.fg |= termbox.ColorRed
	case 32:
		s.fg |= termbox.ColorGreen
	case 33:
		s.fg |= termbox.ColorYellow
	case 34:
		s.fg |= termbox.ColorBlue
	case 35:
		s.fg |= termbox.ColorMagenta
	case 36:
		s.fg |= termbox.ColorCyan
	case 37:
		s.fg |= termbox.ColorWhite
	}
}

func (s *Screen) cursorHandler(command byte, params []int) {
	n, m := 1, 1
	if len(params) > 0 {
		n = params[0]
	}

	if len(params) > 1 {
		m = params[1]
	}

	switch command {
	case ansi.CursorPosition:
		s.x = m - 1
		s.y = n - 1
	case ansi.CursorUp:
		s.y = s.y - n
		if s.y < 0 {
			s.y = 0
		}
	case ansi.CursorDown:
		s.y = s.y + n
		if s.y > s.height {
			s.y = s.height
		}
	case ansi.CursorForward:
		s.x = s.x + n
		if s.x > s.width {
			s.x = s.width
		}
	case ansi.CursorBack:
		s.x = s.x - n
		if s.x < 0 {
			s.x = 0
		}
	}
}

func (s *Screen) Write(line string) {
	defer termbox.Flush()
	s.buffer = append(s.buffer, line)
	lines := []string{line}

	if s.y > s.height {
		lines = s.reset()
	}
	s.writeLines(lines)
}

func (s *Screen) BufferedWriter(screenChan chan []byte) {
	for {
		select {
		case response := <-screenChan:
			s.Write(string(response))
		}
	}
}

func (s *Screen) Resize(width, height int) {
	defer termbox.Flush()
	s.width, s.height = width, height
	s.width--
	s.height--
	lines := s.reset()
	s.writeLines(lines)
}

func (s *Screen) reset() []string {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.x = 0
	s.y = 0
	lines := s.buffer
	if len(s.buffer) > s.height {
		scope := (len(s.buffer) - s.height) + 1
		lines = s.buffer[scope:]
	}
	return lines
}
