package main

import (
	"fmt"
	"runtime"
)

func main(){
	go func(){
		for i:=0;i <5;i++{
			fmt.Println("go")
		}
	}()

	for i:=0;i<2;i++{
		//让出时间片，先让别的线程执行
		runtime.Gosched()
		fmt.Println("hello")
	}
}
