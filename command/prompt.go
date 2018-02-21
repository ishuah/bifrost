package command

import (
	termbox "github.com/nsf/termbox-go"
)

type Prompt struct {
	input                    []rune
	x, y                     int
	cursorIndex, screenWidth int
}

func NewPrompt(screenWidth int) Prompt {
	return Prompt{screenWidth: screenWidth}
}
func (p *Prompt) GetPosition() (int, int) {
	return p.x, p.y
}

func (p *Prompt) ResetCursor() {
	p.draw()
}

func (p *Prompt) MoveCursor(move int) {
	newPos := (p.cursorIndex + move) + p.x
	if newPos < p.x || newPos > len(p.input)+p.x {
		return
	}

	p.cursorIndex = p.cursorIndex + move
	p.draw()
}

func (p *Prompt) Reset() (int, int) {
	p.x = 0
	p.y = 0
	return p.x, p.y
}

func (p *Prompt) Write(line string) {
	for _, char := range line {
		if p.x > p.screenWidth {
			p.x = 0
			p.y++
		}
		// new line character
		if char == 10 {
			p.x = 0
			p.y++
			continue
		}
		termbox.SetCell(p.x, p.y, char, termbox.ColorDefault, termbox.ColorDefault)
		p.x++
	}
}

func (p *Prompt) draw() {
	defer termbox.Flush()
	// clear the the command prompt
	for x := 0; x <= len(p.input)+1; x++ {
		termbox.SetCell(x+p.x, p.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	// draw the input
	for i, char := range p.input {
		termbox.SetCell(i+p.x, p.y, char, termbox.ColorDefault, termbox.ColorDefault)
	}
	// draw the cursor
	termbox.SetCell(p.cursorIndex+p.x, p.y, ' ', termbox.ColorBlack, termbox.ColorWhite)
}

func (p *Prompt) InsertInputChar(char rune) {
	p.input = append(p.input, ' ')
	copy(p.input[p.cursorIndex+1:], p.input[p.cursorIndex:])
	p.input[p.cursorIndex] = char
	p.cursorIndex++
	p.draw()
}

func (p *Prompt) DeleteInputChar() {
	if len(p.input) == 0 {
		return
	}
	copy(p.input[p.cursorIndex-1:], p.input[p.cursorIndex:])
	p.input[len(p.input)-1] = ' '
	p.input = p.input[:len(p.input)-1]
	p.cursorIndex--
	p.draw()
}

func (p *Prompt) ReturnInput() []rune {
	input := append(p.input, '\r')
	p.input = p.input[:0]
	p.cursorIndex = 0
	return input
}

func (p *Prompt) ClearInput() []rune {
	termbox.SetCell(p.cursorIndex+p.x, p.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	input := p.input
	p.input = p.input[:0]
	p.cursorIndex = 0
	return input
}
