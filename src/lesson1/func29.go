package lesson1

import "fmt"

func main(){
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice2(numbers)

	fmt.Println("numbers==" ,numbers)
	fmt.Println("numbers[1:4] ==,",numbers[1:4])
	fmt.Println("numbers[:3]==",numbers[:3])
	fmt.Println("numbers[4:]==",numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice2(numbers1)

	numbers2 := numbers[:2]
	printSlice2(numbers2)

	numbers3 := numbers[2:5]
	printSlice2(numbers3)
}


func printSlice2(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}