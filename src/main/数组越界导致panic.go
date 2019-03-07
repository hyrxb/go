package main

import "fmt"

func testa1(){
	fmt.Println("assss")
}


func testb1(x int){
	var a [10]int
	a[x] =111
}

func testc1(){
	fmt.Println("cccc")
}


func main(){
	testa1()
	testb1(20)
	testc1()
}