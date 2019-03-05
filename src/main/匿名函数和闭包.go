package main

import "fmt"

func main(){
	a :=10
	str := "mike"

	f1 := func(){
		fmt.Println("a=",a)
		fmt.Println("str=",str)
	}
	f1()


	//定义匿名函数同时调用
	func(){
		fmt.Printf("a=%d,str=%s\n",a,str)
	}()


	func(i,j int){
		fmt.Printf("i=%d,j=%d\n",i,j)
	}(10,20)


	x,y := func(i,j int)(min,max int){
		if i> j {
			max = i
			min =j
		}else{
			max = j
			min = i
		}
		return
	}(10,20)

	fmt.Println("x=",x,"y=",y)

}
