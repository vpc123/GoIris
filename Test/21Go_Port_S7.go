package main

import (
	"fmt"
)

func list_elem(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

}

func main() {
	go list_elem(10)
	var input string
	fmt.Scanln(&input)

}
