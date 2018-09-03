package main

import (
	"fmt"
)

type Phone struct {
	price int
	color string
}

func (phone Phone) ringing() {

	fmt.Println("Phone is ringing.....")

}

type IPhone struct {
	Phone
	model string
}

func main() {
	var p IPhone
	p.price = 5000
	p.color = "Black"
	p.model = "iPhone 6"
	fmt.Println("I have a iPhone！")
	fmt.Println("price:", p.price)
	fmt.Println("color:", p.color)
	fmt.Println("model：", p.model)
	p.ringing()

}
