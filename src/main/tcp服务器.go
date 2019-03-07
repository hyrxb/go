package main

import (
	"fmt"
	"net"
)

func main(){
	listener,err := net.Listen("tcp","127.0.0.1:8888")

	if err !=nil{
		fmt.Println("err=",err)
	}

	defer listener.Close()

	//阻塞等待客户端
	for{
		conn,err := listener.Accept()
		if err !=nil{
			fmt.Println("err=",err)
			continue
		}
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("err=",err)
			return
		}
		fmt.Println("buf=",string(buf[:n]))

		defer conn.Close()
	}

}
