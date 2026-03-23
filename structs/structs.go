package structs

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base float64
	Hypotenuse float64
}

func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func (t Triangle) Perimeter() float64{
	return t.Base + t.Height + t.Hypotenuse
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func (c Circle) Area() float64{
	return math.Pi * (math.Pow(c.Radius, 2))
}

func (t Triangle) Area() float64{
	s := t.Perimeter() / 2
	return math.Sqrt((s*(s-t.Height)*(s-t.Hypotenuse)*(s-t.Base)))
}

type Shape interface {
	Area() float64
	Perimeter() float64
}