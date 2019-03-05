package main

import "fmt"

func myfunc(tmp ...int){
	for _,data := range tmp{
		fmt.Println("data = ",data)
	}
}

func myfunc02(tmp ...int){
	for _,data := range tmp{
		fmt.Println("data = ",data)
	}
}


func test01(args ...int){
	//全部元素传递给myfunc
	//myfunc(args...)
	//只想把最后２个参数传递给另外一个函数使用
	//myfunc02(args[:2]...) //args[0] ~args[2] 不包括２
	myfunc02(args[2:]...)

}

func main(){
	test01(1,2,3)
}
