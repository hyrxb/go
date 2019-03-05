package main

import "fmt"

func main(){
	var p *int

	 p = new(int) // p是　*int ，指向int 类型

	 *p = 666

	 fmt.Println("p= ",*p)
}
