package main

import "fmt"

func main(){
	a :=20
	b :=20

	defer func(a,b int){
		fmt.Printf("a=%d,b=%d\n",a,b)
	}(a,b)

	a =11
	b =222
	fmt.Printf("外部 a=%d,b=%d\n",a,b)
}

func main001(){
	a :=20
	b :=20

	defer func(){
		fmt.Printf("a=%d,b=%d\n",a,b)
	}()

	a =11
	b =222
	fmt.Printf("外部 a=%d,b=%d\n",a,b)
}
