package main

import (
	"fmt"
	"net"
	"strings"
)

/**
　* 处理用户请求
 */
func HandleConn(conn net.Conn){
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect_successful :",addr)

	buf := make([]byte,2040)

	for{
		n,err := conn.Read(buf)

		if err != nil{
			fmt.Println("err=",err)
			return
		}

		fmt.Println("read buf=",string(buf[:n]))

		if "exit" == string(buf[:n]){
			fmt.Println("addr exit")
		}

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}


func main(){

	listener,err :=net.Listen("tcp","127.0.0.1:8888")

	if err != nil{
		fmt.Println("err=",err)
		return
	}

	defer listener.Close()

	for{
		conn,err := listener.Accept()

		if err != nil{
			fmt.Println("err=",err)
			return
		}

		go HandleConn(conn)
	}

}
