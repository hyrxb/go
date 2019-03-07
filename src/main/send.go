package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string,conn net.Conn){
	f,err :=os.Open(path)
	if err != nil{
		fmt.Println("os Open err=",err)
		return
	}

	defer f.Close()

	//读文件内容，读多少发多少

	buf := make([]byte,1024*4)
	for{
		n,err := f.Read(buf)
		if err != nil{
			if err ==io.EOF{
				fmt.Println("文件发送文笔")
			}else{
				fmt.Println("Read err=",err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}


func main(){
	fmt.Println("请输入要传输的文件:")
	var path string

	fmt.Scan(&path)


	info,err := os.Stat(path)

	if err != nil{
		fmt.Println("os.Stat err=",err)
	}

	conn,err1 := net.Dial("tcp","127.0.0.1:8888")

	if err1 != nil{
		fmt.Println("conn Write err=",err)
		return
	}

	defer conn.Close()

	//给接收方发送文件名
	_,err = conn.Write([]byte(info.Name()))

	if err != nil{
		fmt.Println("conn Write err=",err)
		return
	}

	var n int

	buf := make([]byte,1024)
	n,err = conn.Read(buf)

	if err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}

	//切片转换成字符串
	if string(buf[:n]) == "ok" {
		sendFile(path,conn)
	}

}
