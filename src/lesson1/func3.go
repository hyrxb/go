package lesson1

import "fmt"

func main(){
	var a int =100
	var b int =200
	fmt.Println("交换前ａ的值为:%d\n",a)
	fmt.Println("交换前ｂ的值为:%d\n",b)

	swap1(a,b)

	fmt.Println("交换后a的值为：%d\n",a)
	fmt.Println("交换后b的值为%d\n",b)

}

func swap1(x,y int) int{
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}