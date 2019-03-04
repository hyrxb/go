package main

import "fmt"

func main(){
	var a int
	//变量声明时没有赋值默认值是０
	fmt.Printf("a = ",a)

	a =10
	fmt.Println("a = ",a)

	var b int =10
	b=20
	fmt.Println("b = ",b)

	c :=30

	fmt.Println("c type is %T\n",c)
}