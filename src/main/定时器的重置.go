package main

import (
	"fmt"
	"time"
)

func main(){

	timer := time.NewTimer(3*time.Second)
	timer.Reset(1 * time.Second)
	fmt.Println("ok=")

	<- timer.C

	fmt.Println("时间到")
}
