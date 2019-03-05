package main

import "fmt"

func main(){
	//s := make([]int ,0,1)
	//oldCap := cap(s)

	srcSlice := []int{1,2}
	dstSlice := []int{6,6,6,6,6}

	copy(dstSlice,srcSlice)

	fmt.Println("dstSlice=",dstSlice)


}
