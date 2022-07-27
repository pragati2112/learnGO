package main

import "fmt"

func main() {
	var xyz geometry = rect{
		10, 20,
	}
	fmt.Println(xyz.area())
	fmt.Println(xyz.perimeter())
}

//interface is a collection of methods/behaviours
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	height, width int
}

//we created methods---
func (r rect) area() float64 {
	return float64(r.width * r.height)
}

func (r rect) perimeter() float64 {
	return float64(2*r.width + 2*r.height)
}
