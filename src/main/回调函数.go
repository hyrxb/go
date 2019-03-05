package main

import "fmt"

type FunType2 func(int,int) int


func Calc(a,b int,fTest FunType2)(result int){
	fmt.Println("Calc")
	result = fTest(a,b)
	return
}

func main(){

}
