package containers

import (
	"github.com/B9O2/canvas/pixel"
)

type HStack struct {
	*BaseFrameContainer
}

func (hs *HStack) Draw(width uint, height uint) (pixel.PixelMap, error) {
	width, height, pm, err := hs.DrawBorder(width, height)
	if err != nil {
		return pm, err
	}

	containers := hs.SubContainers()
	baseX := hs.HPadding()
	baseY := hs.VPadding()

	var widthRanges []Range
	for _, container := range containers {
		switch fc := container.(type) {
		case FrameContainer:
			widthRanges = append(widthRanges, fc.Frame().Width())
		default:
			widthRanges = append(widthRanges, ZeroFrame.Width())
		}
	}

	//Calculate
	widths, err := SpaceAllocate(widthRanges, width)
	if err != nil {
		return pm, err
	}

	//Draw
	offset := uint(0)
	for id, container := range containers {
		cpm, err := container.Draw(widths[id], height)
		if err != nil {
			return pm, err
		}
		cpm.Range(func(x, y uint, p *pixel.Pixel) bool {
			x += offset
			pm.SetPixel(baseX+x, baseY+y, *p)
			return true
		}, nil)
		offset += widths[id]
	}

	return pm, nil
}

//hs *HStack

func NewHStack(containers ...Container) *HStack {
	return &HStack{
		NewBaseFrameContainer(0, 0, containers),
	}
}
