package main

import "fmt"

func my_func03() (int,int,int){
	return 1,2,3
}


func my_func04() (a int,b int,c int){
	a,b,c = 1,2,3
	return
}


func main(){
	a,b,c := my_func03()
	fmt.Printf("a=%d,b=%d,c=%d",a,b,c)

	e,f,g := my_func04()
	fmt.Printf("e=%d,f=%d,g=%d",e,f,g)
}
