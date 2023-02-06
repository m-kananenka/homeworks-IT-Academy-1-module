package box

type Rectangle struct {
	A float64
	B float64
}

func (r Rectangle) Sqr() float64 {
	return r.A * r.B
}

func (r Rectangle) Per() float64 {
	return 2 * (r.A + r.B)
}
func NewRectangle(a float64, b float64) Rectangle {
	return Rectangle{
		A: a,
		B: b,
	}
}
