package box

import "math"

type Circle struct {
	R float64
}

func (c Circle) Sqr() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

func (c Circle) Per() float64 {
	return 2 * math.Pi * c.R
}

func NewCircle(a float64) Circle {
	return Circle{
		R: a,
	}
}
