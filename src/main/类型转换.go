package main

import "fmt"

func main(){
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n",flag)

	//fmt.Printf("flag = %d\n",int(flag))
	//cannot convert flag (type bool) to type int
	//bool 类型不能转换成 int

	//整形页不能转成　bool

}
