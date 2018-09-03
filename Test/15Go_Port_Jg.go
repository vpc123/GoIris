package main

import (
	"fmt"
)

type Phone struct {
	price int
	color string
}
type IPhone struct {
	phone Phone
	model string
}

func main() {
	var p IPhone
	p.phone.price = 5000
	p.phone.color = "bloack"
	p.model = "iPhone 5"
	fmt.Println("I have a iPhone：")
	fmt.Println("prince：", p.phone.price)
	fmt.Println("Color:", p.phone.color)
	fmt.Println("Model：", p.model)

}
