package main

import "fmt"

func test03()(sum int){
	for i:=1;i<=100;i++{
		sum += i
	}
	return
}

func test04(num int)(sum int){
	if num ==1{
		return 1
	}
	sum =  num + test04(num -1)
	return
}


func main(){
	num :=test03()
	fmt.Println("num=",num)
	sum := test04(100)
	fmt.Println("sum=",sum)
}
