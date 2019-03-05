package main

import "fmt"

func main(){
	//num :=1

	var num int
	fmt.Printf("请按下楼层：")

	fmt.Scan(&num)

	switch num{
		case 1:
			fmt.Println("按下的是１楼")
		case 2:
			fmt.Println("按下的是２楼")
		case 3:
			fmt.Println("按下的是３楼")

		case 4:

			fmt.Println("按下的是４楼")

		default:
			fmt.Println("按下的是xxx")
	}
}
