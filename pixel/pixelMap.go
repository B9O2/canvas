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

func (pm PixelMap) DrawRepeatLine(
	x, y, length uint,
	p Pixel,
	vertical bool,
) error {
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

// Display 完整显示PixelMap
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

// DisplayTrimmed 自动裁剪空白边缘后显示
func (pm *PixelMap) DisplayTrimmed(space Pixel) error {
	trimmed, err := pm.Trim()
	if err != nil {
		return err
	}
	return trimmed.Display(space)
}

// GetContentBounds 获取有内容区域的边界
// 返回 (minX, minY, maxX, maxY)，如果没有内容返回 (0, 0, 0, 0)
func (pm PixelMap) GetContentBounds() (minX, minY, maxX, maxY uint) {
	pmWidth := pm.Width()
	pmHeight := pm.Height()

	hasContent := false
	minX, minY = pmWidth, pmHeight
	maxX, maxY = 0, 0

	for y := uint(0); y < pmHeight; y++ {
		for x := uint(0); x < pmWidth; x++ {
			if pm[y][x].Char() != 0 {
				hasContent = true
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	if !hasContent {
		return 0, 0, 0, 0
	}

	return minX, minY, maxX, maxY
}

// Trim 自动裁剪掉空白边缘
// 如果整个画布都是空的，返回1x1的空PixelMap
func (pm PixelMap) Trim() (PixelMap, error) {
	minX, minY, maxX, maxY := pm.GetContentBounds()

	if minX == 0 && minY == 0 && maxX == 0 && maxY == 0 {
		emptyMap, err := NewPixelMap(1, 1)
		return emptyMap, err
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	return pm.Crop(minX, minY, width, height)
}

// Crop 从PixelMap中截取指定区域，返回新的PixelMap
func (pm PixelMap) Crop(x, y, width, height uint) (PixelMap, error) {
	pmWidth := pm.Width()
	pmHeight := pm.Height()

	if x >= pmWidth || y >= pmHeight {
		return nil, fmt.Errorf(
			"crop position (%d,%d) out of bounds (%d,%d)",
			x,
			y,
			pmWidth,
			pmHeight,
		)
	}

	if x+width > pmWidth {
		width = pmWidth - x
	}
	if y+height > pmHeight {
		height = pmHeight - y
	}

	cropped := make([][]Pixel, height)
	for i := uint(0); i < height; i++ {
		cropped[i] = make([]Pixel, width)
		copy(cropped[i], pm[y+i][x:x+width])
	}

	return cropped, nil
}

func NewPixelMap(width, height uint) (PixelMap, error) {
	pm := make([][]Pixel, height)
	for i := uint(0); i < height; i++ {
		pm[i] = make([]Pixel, width)
	}
	return pm, nil
}
