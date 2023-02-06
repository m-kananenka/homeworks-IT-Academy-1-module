package main

import (
	"fmt"
	"homework/homework5/box"
)

func main() {

	var f box.Figure
	f = box.Circle{R: 3.1}
	fmt.Println(f.Sqr())
	fmt.Println(f.Per())

	f = box.Square{A: 3.7}
	fmt.Println(f.Sqr())
	fmt.Println(f.Per())

	f = box.Parallelogram{A: 3, B: 5.1, Height: 4.2}
	fmt.Println(f.Sqr())
	fmt.Println(f.Per())

	f = box.Rectangle{A: 2.5, B: 5.3}
	fmt.Println(f.Sqr())
	fmt.Println(f.Per())

	circle := box.Circle{R: 5}
	var d = box.Cylinder{Circle: circle, Height: 4}
	f = d
	fmt.Println(f.Sqr())
	fmt.Println(f.Per())
	fmt.Println(d.Volume())

	slice := make([]box.Figure, 0, 0)
	slice1 := make([]string, 0, 0)

	var collection = box.NewCollection(slice, slice1)

	collection.AddFigures(f)
	cir := box.Circle{R: 3.1}
	collection.AddFigures(cir)
	fmt.Println(collection)

	fmt.Println(collection.SumSqr())
	fmt.Println(collection.SumPer())

}
