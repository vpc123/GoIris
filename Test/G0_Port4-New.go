package main

import (
	"fmt"
)

func set_value(x_ptr *int) {
	*x_ptr = 100
}
func main() {
	x_ptr := new(int)
	set_value(x_ptr)
	// x_ptr指向地址
	fmt.Println(x_ptr)
	// x_ptr本身地址
	fmt.Println(&x_ptr)
	// x_pr指向的地址值
	fmt.Println(*x_ptr)

}
