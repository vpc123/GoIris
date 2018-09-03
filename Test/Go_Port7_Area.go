package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func main() {
	var rect Rect
	rect.width = 100
	rect.length = 200
	fmt.Println("Area is:")
	fmt.Println(rect.width * rect.length)

}
