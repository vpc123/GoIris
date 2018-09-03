package main

import (
	"fmt"
)

func cheange(x *int) {
	*x = 200
}
func main() {
	var x int = 100
	fmt.Println(x)
	cheange(&x)
	fmt.Println(x)
}
