package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int){
	rand.Seed(time.Now().UnixNano())

	var num int

	for{
		num = rand.Intn(10000)
		if num >= 1000{
			break
		}
	}
	*p = num

}

/**
 获取　４位数的每一位数字
 */
func getNum(s []int,num int){
	s[0] = num /1000
	s[1] = num %1000 /100
	s[2] = num %100 / 10
	s[3] = num % 10
}

func main(){
	var randNum int
	CreateNum(&randNum)
	fmt.Println("num=",randNum)
	randSlice := make([]int,4)
	getNum(randSlice,randNum)
	fmt.Println("randSlice=",randSlice)

}


