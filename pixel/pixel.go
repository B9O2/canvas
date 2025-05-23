package pixel

import "github.com/mattn/go-runewidth"

type Pixel struct {
	char  rune
	color any
	width int
}

func (p *Pixel) Char() rune {
	return p.char
}

func (p *Pixel) Width() int {
	return p.width
}

func NewPixel(char rune, color any) Pixel {
	p := Pixel{
		char:  char,
		color: color,
		width: runewidth.RuneWidth(char),
	}
	return p
}

var (
	Space = NewPixel(' ', nil)
	Dot   = NewPixel('.', nil)
)
