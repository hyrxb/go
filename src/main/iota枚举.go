package main

import "fmt"

func main(){
	//iota 常量自动生成器，每一行，自动累加１
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d,b=%d,c=%d\n",a,b,c)//a = 0,b=1,c=2

	const d = iota //iota 遇到const重置为０
	fmt.Printf("d=%d\n",d) //d=0

	//可以致谢一个iota
	const(
		a1 = iota
		b1
		c1
	)

	fmt.Printf("a1=%d,b1=%d,c1=%d\n",a1,b1,c1)
	//a1=0,b1=1,c1=2i=0


	//如果是同一行，值都一样

	const(
		i = iota
		j1,j2,j3=iota,iota,iota
		k = iota
	)

	fmt.Printf("i=%d,j1=%d,j2=%d,j3=%d,k=%d",i,j1,j2,j3,k)
	//a1=0,b1=1,c1=2i=0,j1=1,j2=1,j3=1,k=2
}
