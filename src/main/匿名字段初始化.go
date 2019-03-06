package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}


type Stu struct{
	Person
	id int
	addr string
}


func main(){
	var s1 Stu = Stu{Person{"mike",'n',18},1,"bj"}

	fmt.Println("s1=",s1)


	s2 := Stu{Person{"mike",'m',18},1,"bj"}

	fmt.Println("s2=",s2)

	//指定成员初始化

	s3 := Stu{id:1}

	fmt.Println("s3=",s3)

	s4 := Stu{Person:Person{name:"mike"},id:1}

	fmt.Println("s4=",s4)

	fmt.Println(s4.Person.name,s1.id)


	var s Stu

	s.name = "mi"

	fmt.Println(s.name)
}