package main

import "fmt"

func testa11(){
	fmt.Println("adssss")
}

func testb11(x int){
	defer func(){
		if err := recover();err !=nil{
			fmt.Println(err)
		}
	}()

	var a [10]int
	a[x] =111
}

func testc11(){
	fmt.Println("sssss")
}


func main(){
	testa11()
	testb11(20)
	testc11()

}

