package main

import "fmt"

func main(){

	m := map[int]string{1:"mike",2:"yoyo",3:"go"}

	for key,value := range m{
		fmt.Printf("%d====>%s\n",key,value)
	}

	//判断一个key值是否存在

	//Go语言的map如何判断key是否存在
	//判断方式为value,ok := map[key], ok为true则存在

	value,ok := m[1]

	if ok == true{
		print(value)
	}
}