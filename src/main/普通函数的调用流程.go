package main

import "fmt"

func funcb()( b int){
	b=222
	fmt.Println("func b= ",b)
	return
}

func funca()(a int){
	a = 111

	b := funcb()
	fmt.Println("func a b=",b)

	fmt.Println("funca a =",a)
	return
}

func main(){
	fmt.Println("man func")
	a := funca()
	fmt.Println("main a = ",a)
}
