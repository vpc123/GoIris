package main

import (
	"fmt"
)

func Phone interface {

	call()
	sales() int
}

type NokiaPhone struct{
	price int 
}

func (nokiaPhone NokiaPhone) call(){

    fmt.Println("I am Nokia, I can call you!")	
}
func (nokiaPhone NokiaPhone) sales() int{
	return nokiaPhone.price
	
}
type IPhone struct{
	price int
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone,I can call you!")
	
}
func (iPhone IPhone) sales() int {
	return iPhone.price
	
}


func func main() {
	var phones = [5]Phone{
		NokiaPhone{price:345},
		IPhone{price:5000},
		IPhone{price:3400},
		NokiaPhone{price:450},
		IPhone{price:5000},
	}	
	var totalSales =0
    for _,phone := range phones {
    	totalSales += phones.sales()
    }
    fmt.Println(totalSales)


} 