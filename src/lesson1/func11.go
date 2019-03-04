package lesson1

import "fmt"

var a int =20

func main(){
	var a int = 10
	var b int =20
	var c int =30

	fmt.Printf("main()函数中　a=%d\n",a)
	c= sum(a,b)
	fmt.Println("main()函数中 c= %d\n",c)
}

func sum(a,b int) int{
	fmt.Printf("sum()函数中 a=%d\n",a)
	fmt.Printf("sum()函数中b=%d\n",b)
	return a + b
}
