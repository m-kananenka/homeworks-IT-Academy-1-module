package box

type Parallelogram struct {
	A      float64
	B      float64
	Height float64
}

func (p Parallelogram) Sqr() float64 {
	return p.A * p.Height
}

func (p Parallelogram) Per() float64 {
	return 2 * (p.A + p.B)
}

func NewParallelogram(a float64, b float64, h float64) Parallelogram {
	return Parallelogram{
		A:      a,
		B:      b,
		Height: h,
	}
}
