package main

import "fmt"

type Myint int

func (a Myint) Add(b Myint) Myint{
	return a + b
}

func Add1(a,b Myint) Myint{
	return a + b
}


func main(){
	var a Myint = 1
	var b Myint = 1

	fmt.Println("a.Add(b) = ",a.Add(b))

	fmt.Println("Add(a,b) = ",Add1(a,b))
}