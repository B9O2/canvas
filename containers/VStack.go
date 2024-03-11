package containers

import (
	"github.com/B9O2/canvas/pixel"
)

type VStack struct {
	*BaseFrameContainer
}

func (vs *VStack) Draw(width, height uint) (pixel.PixelMap, error) {
	width, height, pm, err := vs.DrawBorder(width, height)
	if err != nil {
		return pm, err
	}

	containers := vs.SubContainers()
	baseX := vs.HPadding()
	baseY := vs.VPadding()

	var heightRanges []Range
	for _, container := range containers {
		switch fc := container.(type) {
		case FrameContainer:
			heightRanges = append(heightRanges, fc.Frame().Height())
		default:
			heightRanges = append(heightRanges, ZeroFrame.Height())
		}
	}

	//Calculate
	heights, err := SpaceAllocate(heightRanges, height)
	if err != nil {
		return pm, err
	}

	//Draw
	offset := uint(0)
	for id, container := range containers {
		cpm, err := container.Draw(width, heights[id])
		if err != nil {
			return pm, err
		}
		cpm.Range(func(x, y uint, p *pixel.Pixel) bool {
			y += offset
			pm.SetPixel(baseX+x, baseY+y, *p)
			return true
		}, nil)
		offset += heights[id]
	}

	return pm, nil
}

func NewVStack(containers ...Container) *VStack {
	return &VStack{
		NewBaseFrameContainer(0, 0, containers),
	}
}
