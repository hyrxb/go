package main

import (
	"fmt"
	"time"
)

func main(){

	timer := time.NewTimer(3 * time.Second)

	go func() {
		<- timer.C

		fmt.Println("子携程")
	}()

	timer.Stop()

	for{

	}
}
