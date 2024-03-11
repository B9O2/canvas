package containers

import (
	"github.com/B9O2/canvas/pixel"
)

type Container interface {
	Append(containers ...Container)
	SetBorder(pixel.Pixel)
	DrawBorder(width, height uint) (uint, uint, pixel.PixelMap, error)
	SubContainers() []Container
	Draw(width, height uint) (pixel.PixelMap, error)
}

type FrameContainer interface {
	Container
	SetFrame(minWidth, maxWidth, minHeight, maxHeight uint) error
	Frame() Frame
	SetPadding(pixel.Pixel)
	SetVPadding(padding uint)
	SetHPadding(padding uint)
	VPadding() uint
	HPadding() uint
}

////////////////////////BaseContainer////////////////////

type BaseContainer struct {
	border        pixel.Pixel
	subContainers []Container
}

func (bc *BaseContainer) Append(containers ...Container) {
	bc.subContainers = append(bc.subContainers, containers...)
}

func (bc *BaseContainer) DrawBorder(width, height uint) (uint, uint, pixel.PixelMap, error) {
	pm, err := pixel.NewPixelMap(width, height)
	if err != nil {
		return width, height, pm, err
	}

	if err := pm.DrawRectangle(0, 0, width, height, bc.border); err != nil {
		return width, height, pm, err
	}

	return width, height, pm, nil
}

func (bc *BaseContainer) SetBorder(border pixel.Pixel) {
	bc.border = border
}

func (bc *BaseContainer) SubContainers() []Container {
	return bc.subContainers
}

func NewBaseContainer(containers []Container) *BaseContainer {
	return &BaseContainer{
		subContainers: containers,
	}
}

////////////////////////BaseFrameContainer////////////////////

type BaseFrameContainer struct {
	*BaseContainer
	frame              Frame
	padding            pixel.Pixel
	hPadding, vPadding uint
}

func (bc *BaseFrameContainer) SetFrame(minWidth uint, maxWidth uint, minHeight uint, maxHeight uint) error {
	bc.frame.width.min = minWidth
	bc.frame.width.max = maxWidth

	bc.frame.height.min = minHeight
	bc.frame.height.max = maxHeight
	return nil
}

func (bc *BaseFrameContainer) Frame() Frame {
	return bc.frame
}

func (bc *BaseFrameContainer) SetHPadding(padding uint) {
	bc.hPadding = padding
}

func (bc *BaseFrameContainer) SetVPadding(padding uint) {
	bc.vPadding = padding
}

func (bc *BaseFrameContainer) HPadding() uint {
	return bc.hPadding
}

func (bc *BaseFrameContainer) VPadding() uint {
	return bc.vPadding
}

func (bc *BaseFrameContainer) SetPadding(padding pixel.Pixel) {
	bc.padding = padding
}

// DrawBorder FrameContainer的DrawBorder还会绘制Padding
func (bc *BaseFrameContainer) DrawBorder(width, height uint) (uint, uint, pixel.PixelMap, error) {
	width, height, pm, err := bc.BaseContainer.DrawBorder(width, height)
	if err != nil {
		return width, height, pm, err
	}
	if 2*bc.hPadding > width {
		width = 1
	} else {
		width = width - 2*bc.hPadding
	}

	if 2*bc.vPadding > height {
		height = 1
	} else {
		height = height - 2*bc.vPadding
	}

	err = pm.DrawRectangle(bc.hPadding, bc.vPadding, width, height, bc.padding)
	return width, height, pm, err
}

func NewBaseFrameContainer(hPadding, vPadding uint, containers []Container) *BaseFrameContainer {
	return &BaseFrameContainer{
		BaseContainer: NewBaseContainer(containers),
		hPadding:      hPadding,
		vPadding:      vPadding,
	}
}

// type Rectangle struct {
// 	*BaseContainer
// 	lines, corners [4]byte
// }

// func (r *Rectangle) Draw(width, height uint) (PixelMap, error) {
// 	for _, container := range r.SubContainers() {
// 		container.Draw()
// 	}
// }
