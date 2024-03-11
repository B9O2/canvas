package pixel

type Pixel struct {
	char  rune
	color any
}

func (p *Pixel) Char() rune {
	return p.char
}

func NewPixel(char rune, color any) Pixel {
	p := Pixel{
		char:  char,
		color: color,
	}
	return p
}

var (
	Space = NewPixel(' ', nil)
	Dot   = NewPixel('.', nil)
)
