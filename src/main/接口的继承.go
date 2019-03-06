package main

import "fmt"

type Humaner1 interface{
	sayhi()
}

type Person22 interface{
	Humaner1
	sing(lrc string)
}

type Stu33 struct{
	name string
	id int
}

func (tmp *Stu33) sayhi(){
	fmt.Printf("Stu33[%s,%s]",tmp.name,tmp.id)
}

func (tmp *Stu33) sing(lrc string){
	fmt.Println("Student 在唱歌",lrc)
}


func main(){
	var i Person22

	i = &Stu33{"ss",11}

	i.sayhi()
	i.sing("国歌")
}
