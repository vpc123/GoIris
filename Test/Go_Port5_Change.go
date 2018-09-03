package main

import (
	"fmt"
)

func swap(x, y *int) {
	*x, *y = *y, *x

}

func main() {
	x_val := 100
	y_val := 200
	swap(&x_val, &y_val)
	fmt.Println(x_val)
	fmt.Println(y_val)

}
