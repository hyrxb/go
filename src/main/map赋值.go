package main

import "fmt"

func main(){
	m1 := map[int]string{1:"mike",2:"yoyo"}
	fmt.Println("m1=",m1)

	m1[1] = "c++"
	m1[3] = "go"

	fmt.Println("m1=",m1)

}
