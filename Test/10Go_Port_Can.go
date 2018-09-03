package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func double_area(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 3
	return rect.width * rect.length
}

func main() {
	var rect = Rect{12, 45}
	fmt.Println(double_area(rect))
	fmt.Println("width:", rect.width, "Length:", rect.length)
}
