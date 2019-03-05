package main

import "fmt"

func test02(a int){
	if a ==1{
		fmt.Println("a = ",a)
		return
	}
	test02(a - 1)
	fmt.Println("a = ",a)
}

func main(){
	test02(3)
	fmt.Println("main")
}
