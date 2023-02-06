package box

import "math"

type Square struct {
	A float64
}

func (s Square) Sqr() float64 {
	return math.Pow(s.A, 2)
}

func (s Square) Per() float64 {
	return 4 * s.A
}

func NewSquare(a float64) Square {
	return Square{
		A: a,
	}
}
