package main

import "fmt"

func main(){

	var m1 map[int]string
	fmt.Println("m1=",m1)
	//对于map只有len，没有cap

	fmt.Println("len=",len(m1))

	m2 := make(map[int]string)

	fmt.Println("m2=",m2)
	fmt.Println("len=",len(m2))

	/**
	m3= map[]
l	len= 0
	只是指定了容量，里面一个内容都没有
	 */
	m3 := make(map[int]string,2)
	fmt.Println("m3=",m3)
	fmt.Println("len=",len(m3))


	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"

	fmt.Println("m3=",m3)


	m4 := map[int]string{1:"mike",2:"go",3:"c++"}
	fmt.Println("m4=",m4)

}