package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func main() {
	var rect = Rect{12, 34}
	fmt.Println("Width:", rect.width, "*Length:", rect.length, "=Area:", rect.width*rect.length)
}
