package main

import "fmt"

func main(){
	var t complex128
	t = 2.1 + 3.14i

	fmt.Println("t=",t)

	//自动推导类型

	t2 := 3.3 + 4.4i
	fmt.Println("te type is %T\n",t2)


}