package main

import "fmt"

func main(){
	str := "abc"
	//通过for 打印每个字符

	for i:=0; i< len(str);i++{
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}

	//迭代打印每个元素，默认返回２个值

	for i,data := range str{
		fmt.Printf("str[%d]=%c\n",i,data)
	}
}
