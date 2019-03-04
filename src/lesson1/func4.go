package lesson1

import "fmt"

func main(){
	var a int =100
	var b int =200

	fmt.Println("交换前，ａ的值为:%d\n",a)
	fmt.Println("交换后,b的值为:%d\n",b)

	swap2(&a,&b)

	fmt.Println("交换后，a的值为：%d\n",a)
	fmt.Println("交换后,b的值为:%d\n",b)
}

func swap2(x *int,y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
