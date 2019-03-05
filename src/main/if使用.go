package main

import "fmt"

func main(){

	//s := "王思聪"
	s := "小甜甜"

	if s == "王思聪"{
		fmt.Println("左手一个妹子，右手一个大妈")
	}

	// if 支持一个初始化语句，初始化语句和判断条件以分号分割

	if a:=10;a==10{
		fmt.Println("a == 10")
	}

	b := 11

	if b == 10{
		fmt.Println("b ==10")
	}else{
		fmt.Println("b != 10")
	}

	c := 10

	if(c ==10){
		fmt.Println("c==10")
	}else if c > 10{
		fmt.Println(" c > 10")
	}else if c < 10{
		fmt.Println("c < 10")
	}
}
