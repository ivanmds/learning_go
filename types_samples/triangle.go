package simpleshape

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base float64, height float64) *Triangle {
	return &Triangle{base: base, height: height} //can be used too => &Triangle{base, height}
}

func (t *Triangle) Area() float64 {
	return (t.base / t.height) / 2
}
