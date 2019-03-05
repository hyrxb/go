package main

import "fmt"

func main(){
	//defer 延迟调用,main函数结束前调用
	defer fmt.Println("bbbbb")
	fmt.Println("aaaaa")


}
