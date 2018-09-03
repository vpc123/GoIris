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
	var rect = Rect{100, 200}
	fmt.Print("width:", rect.width, "length:", rect.length, "   Area:", rect.area())
}
