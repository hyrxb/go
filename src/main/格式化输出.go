package main

import "fmt"

func main(){
	a := 10
	b := "abc"
	c := 'a'
	d :=3.14

	fmt.Printf("%T ,%T,%T,%T\n",a,b,c,d)
	fmt.Printf("a=%d,b=%s,c=%c,d=%f\n",a,b,c,d)
	// %v 自动匹配格式输出

	fmt.Printf("a=%v,b=%v,c=%v,d=%v",a,b,c,d)

	/**
	int ,string,int32,float64
	a=10,b=abc,c=a,d=3.140000
	a=10,b=abc,c=97,d=3.14
	 */
}