package main

import "fmt"

func main(){

	switch num :=1;num{
		case 1:
			fmt.Println("按下的是１楼")
		case 2:
			fmt.Println("按下的是２楼")
		case 3:
			fmt.Println("按下的是三楼")
		default:
			fmt.Println("按下的是xx楼")
	}
}
