package main

import "fmt"

func test1111(){
	defer fmt.Println("aaaa")
	fmt.Println("dddddd")
}


func main(){
	go func(){
		fmt.Println("sssss")
		test1111()
		fmt.Println("nnnnnnnn")
	}()

}
