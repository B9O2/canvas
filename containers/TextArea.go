package containers

import (
	"github.com/B9O2/canvas/pixel"
	"github.com/mattn/go-runewidth"
)

type TextArea struct {
	*BaseContainer
	text       string
	color      any
	aliginLeft bool
}

func (ta *TextArea) SetAliginLeft(enabled bool) {
	ta.aliginLeft = enabled
}

func (ta *TextArea) Draw(width, height uint) (pixel.PixelMap, error) {
	_, _, pm, err := ta.DrawBorder(width, height)
	if err != nil {
		return pm, err
	}

	parts := SplitWithLength(ta.text, width)
	lines := uint(len(parts))
	y := uint(0)
	if lines > height {
		parts = parts[:height]
		y = 0
	} else {
		y = (height - lines) / 2
	}
	if lines > 1 || ta.aliginLeft {
		for i, part := range parts {
			pm.DrawLine(0, y+uint(i), PixelString(part, ta.color), false)
		}
	} else if lines == 1 {
		x := (width - uint(runewidth.StringWidth(parts[0]))) / 2
		err = pm.DrawLine(x, y, PixelString(parts[0], ta.color), false)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func NewTextArea(text string) *TextArea {
	return &TextArea{
		BaseContainer: NewBaseContainer(nil),
		text:          text,
	}
}
