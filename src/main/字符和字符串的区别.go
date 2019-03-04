package main

import "fmt"

func main(){
	var ch byte
	var str string

	//字符
	//单引号
	//字符，往往只有一个字符，转义字符除外 '\n'
	ch = 'a'

	//字符串
	//双引号
	//字符串有１个或多个自符
	//字符串都是隐藏了一个结束符号 '\0'
	str = "a"

	fmt.Println("str = ",str)
	fmt.Println("ch = ",ch)
}