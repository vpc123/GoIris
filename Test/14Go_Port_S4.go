package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func (rect *Rect) double_area() float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}
func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200
	fmt.Println(*rect)
	fmt.Println(rect.double_area())
	fmt.Println(*rect)
}
