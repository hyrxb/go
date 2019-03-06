package main

import "fmt"

type Student5 struct{
	id int
	name string
	sex byte
	age int
	addr string
}


func main(){
	s := Student5{1,"mike",'m',18,"bj"}

	fmt.Println("s=",s)
}
