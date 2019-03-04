package lesson1

import "fmt"

func main(){
	var numbers []int
	PrintSlice3(numbers)

	numbers = append(numbers,0)
	PrintSlice3(numbers)

	numbers = append(numbers,1)

	PrintSlice3(numbers)

	numbers = append(numbers,2,3,4)
	PrintSlice3(numbers)

	numbers1 := make([]int ,len(numbers),(cap(numbers))*2)

	copy(numbers1,numbers)

	PrintSlice3(numbers1)


}

func PrintSlice3(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
