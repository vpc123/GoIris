package main

import (
	"fmt"
)

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am Nokia,I can call you！")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am Iphone,I can call you！")
}
func main() {
	var nokia NokiaPhone
	nokia.call()
	var iPhone IPhone

	iPhone.call()
}
