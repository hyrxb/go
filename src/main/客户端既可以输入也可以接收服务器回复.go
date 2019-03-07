package main

import (
	"fmt"
	"net"
	"os"
)

func main(){

	conn,err := net.Dial("tcp","127.0.0.1:8888")

	if err != nil{
		fmt.Println("net Dial err = ",err)
		return
	}

	defer conn.Close()


	//接收服务器回复的数据

	go func(){

		buf := make([]byte,1024)

		for{
			n,err := conn.Read(buf)

			if err != nil{
				fmt.Println("conn Read err=",err)
				return
			}

			fmt.Println(buf[:n]) //打印接收的内容
		}
	}()

	for{
		str := make([]byte,1024)
		n,err := os.Stdin.Read(str)

		if err !=nil{
			fmt.Println("os.Stdin.err=",err)
			return
		}

		fmt.Println(str[:n])
		conn.Write(str[:n])
	}
}
