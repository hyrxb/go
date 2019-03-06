package main

import "fmt"

type Student struct{
	id int
	name string
	sex byte
	age int
	addr string
}


func main(){
	var s1 Student = Student{1,"mike",1,18,"bj"}
	fmt.Println("s1=",s1)

	//指定成员初始化,没有初始化的成员，自动赋值为0

	s2 := Student{name:"mike",addr:"tianj"}
	fmt.Println("s2=",s2)
}