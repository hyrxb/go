package main

import "fmt"

func MaxAndMin(a ,b int)(max,min int){
	if a > b{
		max = a
		min = b
	}else{
		max = b
		min = a
	}
	return max,min
}

func main(){
	max,min := MaxAndMin(10,30)
	fmt.Printf("max=%d,min=%d",max,min)
}
