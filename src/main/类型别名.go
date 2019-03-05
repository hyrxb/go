package main

import "fmt"

func main(){

	type bigint int64

	var a bigint

	fmt.Printf("a type is %T\n",a)


	type(
		long int64
		char byte
	)

	var b long =11
	var c char = 'c'

	fmt.Println("b=",b)
	fmt.Println("c=",c)


}
