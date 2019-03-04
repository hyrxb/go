package main

import "fmt"

func main(){
	var str1 string

	str1 = "abc"

	fmt.Println("str1 = ",str1)

	//自动推导类型
	str2 := "mike"

	fmt.Printf("str2 类型是 %T\n",str2)

	fmt.Println("str2 = ",str2)

	//内建函数，字符串的长度

	fmt.Printf("str2 的长度 %d",len(str2))
}