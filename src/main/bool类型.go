package main

import "fmt"

func main(){
	//声明变量

	var a bool //声明没有初始化值为false
	fmt.Println("a = ",a)
	a  = true

	fmt.Println("a = ",a)


	//自动推导类型

	var b = false

	fmt.Println("b=",b)

	c := false

	fmt.Println("c = ",c)


}
