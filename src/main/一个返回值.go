package main

import "fmt"

//无参有返回值，只有一个返回值
func my_func01() int{
	return 666
}

//给返回值起一个变量名，go推荐写法

func my_func02() (result int){
	result = 666
	return result
}


func main(){
	var a int
	a = my_func01()
	fmt.Println("a = ",a)

	b := my_func01()
	fmt.Println("b = ",b)

	c := my_func02()
	fmt.Println("c = ",c)
}
