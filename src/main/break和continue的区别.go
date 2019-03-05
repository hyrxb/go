package main

import (
	"fmt"
	"time"
)


func main(){
	i :=0

	for{  // for后面不写任何东西，这个循环永远为真，死循环
		i++
		time.Sleep(time.Second)
		if i== 5 {
			break
		}
		fmt.Println("i=",i)
	}
}
