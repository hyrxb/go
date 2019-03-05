package main

import "fmt"

var a int

func test008(){
	fmt.Println("test a = ",a)
}
func main(){
	a = 10
	fmt.Println("a = ",a)

	test008()
}
