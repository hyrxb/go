package main

import "fmt"

func main(){
	 a := []int{0,1,2,3,4,5,6,7,8,9}

	 s1 := a[2:5]

	 s1[0] =666

	 fmt.Println("s1=",s1)

	 s2 := a[:]

	 fmt.Println("s2=",s2)

	 fmt.Printf("len=%d,cap=%d",len(s2),cap(s2))
}
