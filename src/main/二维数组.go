package main

import "fmt"

func aa(args ...int){
	fmt.Println(args)
}

func main(){

	var a[3][4]int

	k:=0

	for i:= 0;i<3;i++{
		for j:= 0;j<4;j++{
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]=%d\t",i,j,a[i][j])
		}
		fmt.Printf("\n")
	}

	aa()
}


