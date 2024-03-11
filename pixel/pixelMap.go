package pixel

import (
	"errors"
	"fmt"
)

type PixelMap [][]Pixel

// Range 执行后会写入PixelMap
func (pm PixelMap) Range(f, endl func(x, y uint, p *Pixel) bool) {
	for y, line := range pm {
		var x int
		var pixel Pixel
		for x, pixel = range line {
			goon := f(uint(x), uint(y), &pixel)
			pm[y][x] = pixel
			if !goon {
				return
			}
		}
		if endl != nil {
			endl(uint(x), uint(y), &pixel)
		}
		pm[y][x] = pixel
	}
}

func (pm PixelMap) SetPixel(x, y uint, p Pixel) (err error) {
	//fmt.Printf("(%d,%d) %s\n", x, y, string(p.Char()))
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()
	pm[y][x] = p
	return nil
}

func (pm PixelMap) Width() uint {
	return uint(len(pm[0]))
}

func (pm PixelMap) Height() uint {
	return uint(len(pm))
}

func (pm PixelMap) DrawLine(x, y uint, pixels []Pixel, vertical bool) error {
	if !vertical {
		for i, p := range pixels {
			//fmt.Printf("(%d,%d) %s\n", x+uint(i), y, string(p.Char()))
			pm.SetPixel(x+uint(i), y, p)
		}
	} else {
		for i, p := range pixels {
			pm.SetPixel(x, y+uint(i), p)
		}
	}

	return nil
}

func (pm PixelMap) DrawRepeatLine(x, y, length uint, p Pixel, vertical bool) error {
	if !vertical {
		for i := uint(0); i < length; i++ {
			pm.SetPixel(x+i, y, p)
		}
	} else {
		for i := uint(0); i < length; i++ {
			pm.SetPixel(x, y+i, p)
		}
	}

	return nil
}

func (pm PixelMap) DrawRectangle(x, y, width, height uint, p Pixel) error {
	//fmt.Println(width, "x", height)
	if err := pm.DrawRepeatLine(x, y, width, p, false); err != nil {
		return err
	}
	if err := pm.DrawRepeatLine(x+width-1, y, height, p, true); err != nil {
		return err
	}
	if err := pm.DrawRepeatLine(x, y+height-1, width, p, false); err != nil {
		return err
	}
	if err := pm.DrawRepeatLine(x, y, height, p, true); err != nil {
		return err
	}
	return nil
}

func (pm *PixelMap) Display(space Pixel) error {
	pm.Range(func(x, y uint, p *Pixel) bool {
		if p.Char() == 0 {
			fmt.Print(string(space.Char()))
		} else {
			fmt.Print(string(p.Char()))
		}
		return true
	}, func(x, y uint, p *Pixel) bool {
		fmt.Print("\n")
		return true
	})

	return nil
}

func NewPixelMap(width, height uint) (PixelMap, error) {
	pm := make([][]Pixel, height)
	for i := uint(0); i < height; i++ {
		pm[i] = make([]Pixel, width)
	}
	return pm, nil
}
