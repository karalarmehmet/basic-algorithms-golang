package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func main() {

	rect := Rectangle{width: 5, height: 3}

	var s Shape = rect

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
