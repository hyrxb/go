package main

import "fmt"

func testa(){
	fmt.Println("aaaaaa")
}

func testb(){
	fmt.Println("bbbbbbbbb")
	//显示调用panic函数，导致程序终端
	panic("this is a panic test")
}

func testc(){
	fmt.Println("CCCCCCC")
}

func main(){
	testa()
	testb()
	testc()
}