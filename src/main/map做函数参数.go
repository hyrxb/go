package main

import "fmt"


func test21(m map[int]string){
	delete(m,1)
}

func main(){
	m := map[int]string{1:"mike",2:"yoyo",3:"go"}

	fmt.Println("m = ",m)

	test21(m)

	fmt.Println("m = ",m)


}
