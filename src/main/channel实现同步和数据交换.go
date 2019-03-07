package main

import (
	"fmt"
	"time"
)

func main(){

	ch := make(chan string)

	defer fmt.Println("主携程页结束了")

	go func(){
		defer fmt.Println("子携程调用玩吧")

		for i:=0;i<2;i++{
			fmt.Println("子携程 1=",i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程，我工作完毕"
	}()

	str := <-ch //死锁
	fmt.Println("str=",str) //fatal error: all goroutines are asleep - deadlock!
}
