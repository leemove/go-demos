package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

type ColorPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) goRight() {
	p.X = 20 + 20
}

func main() {
	var cp ColorPoint
	cp.X = 20
	cp.Y = 300
	cp.goRight()
	fmt.Print(cp.X)
}
