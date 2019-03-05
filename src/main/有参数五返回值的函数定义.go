package main

import "fmt"

func myFun01(a int){
	a = 111
	fmt.Println("a = ",a)
}


func myFun02(a int , b int){
	fmt.Printf("a = %d,b=%d\n",a,b)
}

func myFunc03(a,s string,c float64,d,e int){
	fmt.Printf("a=%s,s=%s,c=%v,d=%d,e=%v",a,s,c,d,e)
}


func main(){
	//有参数无返回值函数调用；函数名（所需函数）
	//调用函数传递的参数叫实参
	myFun01(666)
	myFun02(666,777)

	myFunc03("a","b",12.2,1,2)
}
