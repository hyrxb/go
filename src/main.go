package main

import (
	"calc"
	"fmt"
)

func main(){
	//test()

	num := calc.Add(1,2)
	fmt.Println("num=",num)

	num1 :=calc.Minus(10,2)
	fmt.Println("num1=",num1)
}