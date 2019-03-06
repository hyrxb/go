package main

import "fmt"

type Person11 struct{
	name string
	sex byte
	age int
}


func (tmp *Person11) PrintInfo(){
	fmt.Printf("name=%s,sex=%c,age=%d",tmp.name,tmp.sex,tmp.age)
}


type Stu11 struct{
	Person11
	id int
	addr string
}

func (tmp *Stu11) PrintInfo(){
	fmt.Println("Stu11=",tmp)
}

func main(){
	s := Stu11{Person11{"na",'m',11},1,"bj"}
	s.PrintInfo()
	s.Person11.PrintInfo()
}