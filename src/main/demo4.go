package main

import "fmt"

func main(){
	//赋值，赋值钱，必须先申明变量
	//a = 10

	var a int
	a = 10
	fmt.Println("a = ",a)

	// := 自动推导类型，先申明变量b,在给b赋值
	b :=20
	fmt.Println("b = ",b)

	//b := 30  //前面已经有变量b,不能在创建一个变量ｂ
}
