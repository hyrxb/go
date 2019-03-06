package main

import "fmt"

type Per struct{
	name string
	sex byte
	age int
}

func (tmp Per) PrintInfo(n string,s byte,a int){
	tmp.name = n
	tmp.sex = s
	tmp.age =a

	fmt.Println("tmp=",tmp)
}

func (tmp *Per) setInfo(n string,s byte,a int){
	tmp.name = n
	tmp.sex = s
	tmp.age =a
	fmt.Println("tmp=",tmp)
}



func main(){

	var s1 Per
	fmt.Println("s1=",s1)

	s1.PrintInfo("mike",'a',10)
	fmt.Println("s1=",s1)

	s1.setInfo("mike",'a',20)

	fmt.Println("s1=",s1)

}
