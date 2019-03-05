package main

import "fmt"

func main(){
	//定义在｛｝里面的变量就是局部变量，只能在{}里面有效

	{
		i :=10
		fmt.Println("i=",i)
	}
	//i=111

	if flag :=3;flag ==3{
		fmt.Println("flag = ",flag)
	}

}
