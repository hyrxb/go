package main

import "fmt"

func main(){
	 //变量，程序运行期间，可以改变的量，变量申明需要var
	 //常量，程序运行期间，不可以改变的量，常量申明需要const

 	const a int =10
 	//a = 20 //cannot assign to a 常量不允许修改
 	fmt.Printf("a=%d\n",a)

 	const b = 10 //常量自动推倒类型，没有使用 :=
 	fmt.Printf("b is type %T\n",b)
 	fmt.Println("b=",b)
}
