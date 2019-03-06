package main

import "fmt"

type mystr string

type Person1 struct{
	name string
	sex byte
	age int
}


type Stu1 struct{
	Person1
	int //基础类型匿名字段
	mystr
}


func main(){
	s := Stu1{Person1{"mike",'m',18},1,"sss"}

	fmt.Println("s = ",s)


}
