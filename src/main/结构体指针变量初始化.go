package main

import "fmt"

type Student1 struct{
	id int
	name string
	sex byte
	age int
	addr string
}

func main(){
	var p1 *Student1 = &Student1{1,"mike",'m',18,"bj"}
	fmt.Println("p1=",p1)

	p2 := &Student1{name:"xiao",addr:"xxx"}

	fmt.Println("p2=",*p2)

}
