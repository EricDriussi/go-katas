package poligon

import "fmt"

func Run(poligon Poligon) float64 {
	area := poligon.Area()
	fmt.Println(area)
	return area
}

type Poligon interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (this Triangle) Area() float64 {
	return (this.Height * this.Base) / 2
}

type Square struct {
	Side float64
}

func (this Square) Area() float64 {
	return this.Side * this.Side
}

type Rectangle struct {
	LongSide  float64
	ShortSide float64
}

func (this Rectangle) Area() float64 {
	return this.ShortSide * this.LongSide
}
