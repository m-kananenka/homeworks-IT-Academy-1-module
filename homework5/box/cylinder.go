package box

import "math"

type Cylinder struct {
	Circle
	Height float64
}

func (c Cylinder) Sqr() float64 {
	return 2 * math.Pi * c.R * (c.Height + c.R)

}

func (c Cylinder) Volume() float64 {
	return 2 * math.Pi * math.Pow(c.R, 2) * c.Height
}

func (c Cylinder) Per() float64 {
	return 2 * (c.R + c.Height)
}

func NewCylinder(с Circle, h float64) Cylinder {
	return Cylinder{
		Circle: с,
		Height: h,
	}
}
