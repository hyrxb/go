package main

import "fmt"

func main(){

	var a int =10
	//每个变量有２层含义，变量的内存，变量的地址
	fmt.Printf("%p\n",a)
	fmt.Printf("%v\n",&a)


	var p *int

	p = &a

	fmt.Printf("p=%v,&a=%v\n",p,&a)


	*p = 666

	fmt.Printf("p=%v,&a=%v\n",p,&a)


}
