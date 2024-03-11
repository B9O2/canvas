package containers

type Range struct {
	min uint
	max uint
}

func (r Range) Min() uint {
	return r.min
}

func (r Range) Max() uint {
	return r.max
}

type Frame struct {
	width  Range
	height Range
}

func (f Frame) Width() Range {
	return f.width
}

func (f Frame) Height() Range {
	return f.height
}

var ZeroFrame = Frame{
	width:  Range{0, 0},
	height: Range{0, 0},
}
