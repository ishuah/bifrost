package command

import termbox "github.com/nsf/termbox-go"

type Prompt struct {
	input               []rune
	x, y                int
	cursor, screenWidth int
}

func NewPrompt(screenWidth int) Prompt {
	return Prompt{screenWidth: screenWidth}
}
func (p *Prompt) GetPosition() (int, int) {
	return p.x, p.y
}

func (p *Prompt) UpdatePosition() {
	p.cursor = p.x
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

func (p *Prompt) Draw() {
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
	termbox.SetCell(p.cursor, p.y, ' ', termbox.ColorBlack, termbox.ColorWhite)
}

func (p *Prompt) InsertInputChar(char rune) {
	index := p.cursor - p.x
	p.input = append(p.input, ' ')
	copy(p.input[index+1:], p.input[index:])
	p.input[index] = char
	p.cursor++
	p.Draw()
}

func (p *Prompt) DeleteInputChar() {
	if len(p.input) == 0 {
		return
	}
	index := (p.cursor - p.x) - 1
	copy(p.input[index:], p.input[index+1:])
	p.input[len(p.input)-1] = ' '
	p.input = p.input[:len(p.input)-1]
	p.cursor--
	p.Draw()
}

func (p *Prompt) ReturnInput() []rune {
	input := append(p.input, '\r')
	p.input = p.input[:0]
	return input
}
