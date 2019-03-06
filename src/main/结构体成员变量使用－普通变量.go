package main

import "fmt"

type Student2 struct{
	id int
	name string
	sex byte
	age int
	addr string
}


func main(){
	var s Student2
	s.id =1
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	fmt.Println("s=",s)
}