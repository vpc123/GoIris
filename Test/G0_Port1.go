package main

import (
	"fmt"
)

func main() {

	var x int
	var x_ptr *int
	x = 10
	x_ptr = &x
	fmt.Println(x)
	fmt.Println(x_ptr)
	fmt.Println(*x_ptr)

}
