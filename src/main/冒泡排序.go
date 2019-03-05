package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)
	for i :=0;i <n;i++{
		a[i] = rand.Intn(100)
		fmt.Printf("%d,",a[i])
	}
	fmt.Printf("\n")
}
