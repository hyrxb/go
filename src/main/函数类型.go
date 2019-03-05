package main

import "fmt"

func Add(a,b int) int{
	return a + b
}

func minus(a,b int) int{
	return a - b
}

//FuncType 它是一个函数类型
type FuncType func(int,int) int

func main(){
	var result int
	result  = Add(1,1)
	fmt.Println("result = ",result)


	var fTest FuncType
	fTest = Add
	result = fTest(10,20)
	fmt.Println("result2 = ",result)
}

