package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand01(){
	/**
	rand= 4343637058903381868
rand= 3769183255805726892
rand= 1923662109321608638
rand= 1818688891928401469
rand= 4144162958715305555
	 */
	rand.Seed(666) //设置种子,如果种子一样每次的随机数都一样

	for i :=0;i<5;i++{
		fmt.Println("rand=",rand.Int())
	}
}

func rand02(){

	rand.Seed(time.Now().UnixNano()) //设置种子,如果种子一样每次的随机数都一样

	for i :=0;i<5;i++{
		fmt.Println("rand=",rand.Int())
	}
}

func main(){
	//rand01()
	rand02()
}
