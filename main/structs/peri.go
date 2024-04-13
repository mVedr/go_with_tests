package structs

import "math"

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// returns area of a rectangle 2*(l+b)
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
