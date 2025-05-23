package containers

import (
	"github.com/B9O2/canvas/pixel"
	"github.com/mattn/go-runewidth"
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
		currentLineWidth := uint(runewidth.StringWidth(line))
		if currentLineWidth > artWidth {
			artWidth = currentLineWidth
		}

		if currentLineWidth > width {
			// Truncate line based on display width
			truncatedLine := ""
			currentTruncatedWidth := uint(0)
			for _, r := range []rune(line) {
				runeW := uint(runewidth.RuneWidth(r))
				if currentTruncatedWidth+runeW > width {
					break
				}
				truncatedLine += string(r)
				currentTruncatedWidth += runeW
			}
			lines[id] = truncatedLine
			// Update artWidth if the original line was the widest and got truncated
			if currentLineWidth == artWidth {
				artWidth = currentTruncatedWidth
			}
		} else {
			lines[id] = line
		}
	}

	// After processing all lines, ensure artWidth does not exceed the container's width
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
