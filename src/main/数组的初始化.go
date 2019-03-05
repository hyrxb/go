package main

import "fmt"

func main(){
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a = ",a)

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	c := [5]int{1,2,3}
	fmt.Println(c)

}
