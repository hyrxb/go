package main

import "fmt"

func main(){

	s1 := []int{1,2,3,5}

	fmt.Println("s1=",s1)

	//借助make函数，格式 make 切片类型，长度，容量
	s2 := make([]int,5,10)

	fmt.Println("len=%d,cap=%d\n",len(s2),cap(s2))
}
