package main

import "fmt"

var a byte //全局变量

func main(){
	var a int //局部变量

	fmt.Printf("%T",a)

	{
		var a float32
		fmt.Printf("%T\n",a)
	}
}
