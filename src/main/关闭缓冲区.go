package main

import "fmt"

func main(){
	ch := make(chan int)

	go func() {
		for i:=0;i<3;i++{
			ch <- i
		}
		close(ch)
		//关闭channel　无法在发送数据
	}()

	for{
		if num,ok := <-ch;ok ==true{
			fmt.Println("num=",num)
		}else{
			break
		}
	}
}
