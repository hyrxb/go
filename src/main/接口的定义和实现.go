package main

import "fmt"

type Humaner interface {
	//方法只有申明没有实现
	sayhi()
}


type St1 struct{
	name string
	id int
}

func (tmp *St1) sayhi(){
	fmt.Printf("Student=",)
}


type Teacher struct{
	addr string
	group string
}

func (tmp *Teacher) sayhi(){
	fmt.Printf("Teacher[%s,%s] sayhi\n",tmp.addr,tmp.group)
}


func main(){
	var i Humaner

	s := &St1{"name1",111}
	i =s
	i.sayhi()

	t := &Teacher{"bj","go"}
	i = t
	i.sayhi()
}