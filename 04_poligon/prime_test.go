package poligon_test

import (
	"katas/04_poligon"
	"testing"
)

func TestSquare(t *testing.T) {
	side := 5.0
	square := poligon.Square{Side: side}
	area := poligon.Run(square)
	if area != side*side {
		t.Fail()
	}
}

func TestRectangle(t *testing.T) {
	longSide := 10.0
	shortSide := 5.0
	rectangle := poligon.Rectangle{LongSide: longSide, ShortSide: shortSide}
	area := poligon.Run(rectangle)
	if area != longSide*shortSide {
		t.Fail()
	}
}

func TestTriangle(t *testing.T) {
	base := 8.0
	height := 5.0
	triangle := poligon.Triangle{Base: base, Height: height}
	area := poligon.Run(triangle)
	if area != (base*height)/2 {
		t.Fail()
	}
}
