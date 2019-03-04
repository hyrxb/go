package main

import "fmt"

func main() {

	var f1 float32
	fmt.Println("f1 = ", f1)

	//自动推导类型

	f2 := 3.14
	fmt.Println("f2 = ", f2)
	fmt.Printf("f2 type is %T", f2)

	//f1 =  0
	//f2 =  3.14
	//f2 type is float64

}
