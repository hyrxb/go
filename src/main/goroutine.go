package main

import (
	"fmt"
	"time"
)

func newTask(){
	for{
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}
}

func main(){
	go newTask()

	for{
		fmt.Println("this is a main gorouine")
		time.Sleep(time.Second)
	}

}
