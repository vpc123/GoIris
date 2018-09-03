package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func main() {
	var rect = Rect{width: 100, length: 300}
	fmt.Println("结构体输出结果:")
	fmt.Println(rect.width * rect.length)
}
