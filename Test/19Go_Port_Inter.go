package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I cll you！")
}

type IPhone struct {
}

func (IPhone IPhone) call() {
	fmt.Println("i am iPhone,i can call you！")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

}
