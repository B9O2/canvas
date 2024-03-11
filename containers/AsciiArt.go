package containers

import (
	"github.com/B9O2/canvas/pixel"
)

type AsciiArt struct {
	*BaseContainer
	lines []string
	color any
}

func (art *AsciiArt) Draw(width, height uint) (pixel.PixelMap, error) {
	//padding diabled
	_, _, pm, err := art.DrawBorder(width, height)
	if err != nil {
		return pm, err
	}

	artHeight := uint(len(art.lines))
	if artHeight <= 0 {
		return pm, nil
	}
	lines := make([]string, len(art.lines))
	if artHeight > height {
		copy(lines, art.lines[:height])
	} else {
		copy(lines, art.lines)
	}

	artWidth := uint(0)
	for id, line := range lines {
		l := uint(len(line))
		if l > artWidth {
			artWidth = l
		}
		if l > width {
			line = line[:width]
		}
		lines[id] = line
	}
	if artWidth > width {
		artWidth = width
	}
	if artHeight > height {
		artHeight = height
	}

	x := (width - artWidth) / 2
	y := (height - artHeight) / 2

	for id, line := range lines {
		//fmt.Printf("(%d,%d) %s", x, y+uint(id), line)
		err = pm.DrawLine(x, y+uint(id), PixelString(line, art.color), false)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func NewAsciiArt(lines []string) *AsciiArt {
	return &AsciiArt{
		BaseContainer: NewBaseContainer(nil),
		lines:         lines,
	}
}
