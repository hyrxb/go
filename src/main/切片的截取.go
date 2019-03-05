package main

import "fmt"

func main(){
	array := []int{0,1,2,3,4,5,6,7,8,9}

	s1 := array[:] //不指定长度和容量一样

	fmt.Println("s1 = ",s1)
	fmt.Printf("len=%d,cap=%d\n",len(s1),cap(s1))


	data := array[1]
	fmt.Println("data = ",data)


	s2 := array[2:6:7]

	fmt.Println("s2=",s2)

	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))
}
