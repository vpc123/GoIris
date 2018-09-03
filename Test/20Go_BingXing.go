package main

import (
	"fmt"
	"time"
)

func get_sum_of_divisible(num int, divider int, resultChan chan int) {
	sum := 0
	for value := 0; value < num; value++ {
		if value%divider == 0 {
			sum += value
		}
	}

	resultChan <- sum

}

func main() {
	LIMIT := 10
	resultChan := make(chan int, 3)
	// resultChan := make(chan int, 3)
	t_start := time.Now()
	go get_sum_of_divisible(LIMIT, 3, resultChan)
	sum3 := <-resultChan
	go get_sum_of_divisible(LIMIT, 5, resultChan)
	sum5 := <-resultChan
	// 单独算被15整除
	go get_sum_of_divisible(LIMIT, 15, resultChan)
	sum15 := <-resultChan
	sum := sum3 + sum5 - sum15
	t_end := time.Now()
	fmt.Println(sum)
	fmt.Println(t_end.Sub(t_start))
}
