package main

import "fmt"

type P1 struct{
	name string
	sex byte
	age int
}

type S1 struct{
	*P1
	id int
	addr string
}


func main(){
	s1 := S1{&P1{"aaa",'m',18},1,"bj"}

	fmt.Println(s1.name,s1.sex,s1.age,s1.id,s1.addr)
}