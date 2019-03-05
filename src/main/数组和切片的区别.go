package main

import "fmt"

func main(){
	a := [5]int{} //数组长度是一个固定的常量，数组不能修改长度，len和cap 永远是5

	fmt.Printf("len=%d,cap=%d\n",len(a),cap(a))

	//切片
	//[] 里面是空，或者...，切片的容量和长度可以不固定
	s := []int{}

	s = append(s,11)

	fmt.Printf("len=%d,cap=%d\n",len(s),cap(s))

}
