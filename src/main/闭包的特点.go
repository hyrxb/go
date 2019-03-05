package main

import "fmt"

func test002() int{
	//函数调用是，ｘ才分配空间
	var x int
	x++
	return x*x //函数调用完毕,x自动释放
}

/**
函数的返回值是一个匿名函数，返回一个函数类型
 */
func test003() func() int{
	var x int //没有初始化，值为0
	return func() int{
		x++
		return x *x
	}
}

func main(){
	f := test003()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}


func main01(){
	fmt.Println(test002())
	fmt.Println(test002())
	fmt.Println(test002())
	fmt.Println(test002())
}
