package main

import "fmt"

type Man struct{
	name string
	sex byte
	age int
}


func (tmp Man) PrintInfo(){
	fmt.Println("tmp=",tmp)
}


func (tmp Man) SetInfo(n string,s byte,a int){
	tmp.name = n
	tmp.sex = s
	tmp.age =a
}


func main(){
	p := Man{"xxx",'s',11}
	p.PrintInfo()

}
