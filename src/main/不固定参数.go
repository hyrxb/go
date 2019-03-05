package main

import "fmt"

func myFunc01(a int,b int){

}

func myFunc05(args ...int){
	fmt.Println("len(args) = ",len(args))

	for i :=0;i<len(args);i++{
		fmt.Printf("args[%d] = %d\n",i,args[i])
	}

	for i,data := range args{
		fmt.Printf("args[%d] = %d\n",i,data)
	}
}

//固定参数一定要传参，不定参数根据需求传递
func myFunc06(a int,args ...int){
	fmt.Println("a = ",a)
}


func main(){
	myFunc05()
	myFunc05(1)
	myFunc05(1,2,3)

	myFunc06(1)
}
