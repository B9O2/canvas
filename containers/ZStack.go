package containers

import (
	"github.com/B9O2/canvas/pixel"
)

type ZStack struct {
	*BaseContainer
}

func (zs *ZStack) Draw(width uint, height uint) (pixel.PixelMap, error) {
	width, height, pm, err := zs.DrawBorder(width, height)
	if err != nil {
		return pm, err
	}

	containers := zs.SubContainers()

	//Draw
	for _, container := range containers {
		cpm, err := container.Draw(width, height)
		if err != nil {
			return pm, err
		}
		cpm.Range(func(x, y uint, p *pixel.Pixel) bool {
			if p.Char() != 0 {
				pm.SetPixel(x, y, *p)
			}
			return true
		}, nil)
	}

	return pm, nil
}

func NewZStack(containers ...Container) *ZStack {
	return &ZStack{
		NewBaseContainer(containers),
	}
}
