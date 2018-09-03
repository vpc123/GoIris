package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200
	fmt.Print(rect.area())
}
