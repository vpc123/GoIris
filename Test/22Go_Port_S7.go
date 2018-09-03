package main

import (
	"fmt"
	"math/rand"
	"time"
)

func list_elem(n int, tag string) {
	for i := 0; i < n; i++ {
		fmt.Println(tag, i)
		tick := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * tick)
	}
}

func main() {
	go list_elem(10, "go_a")
	go list_elem(20, "go_b")
	var input string
	fmt.Scanln(&input)
}
