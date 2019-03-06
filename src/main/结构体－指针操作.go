package main

import "fmt"

type Student3 struct{
	id int
	name string
	sex byte
	age int
	addr string
}


func main(){
	var s Student3

	//定义一个指针变量，保存s的地址

	var p1 *Student3
	p1 = &s

	p1.id = 18
	(*p1).name = "mike"
	p1.sex = 'n'
	p1.age= 18
	p1.addr = "bj"

	fmt.Println("p1=",p1)

	//通过new 申请一个结构体
	p2 := new(Student3)

	p2.id = 2
	p2.name="mike"
	p2.sex = 'm'
	p2.addr = "nj"
	fmt.Println("p1=",p1)
	fmt.Println("p2=",p2)
}
