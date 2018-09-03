package main

import (
	"fmt"
)

type Phone struct {
	price int
	color string
}

type IPhone struct {
	Phone
	model string
}

func main() {
	var p IPhone
	p.price = 5000
	p.color = "Block"
	p.model = "IPhone5"
	fmt.Print("i have a phone")
	fmt.Println("price:", p.price)
	fmt.Print("Model：", p.model)

}
