package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println(strings.Contains("hellogo","hello"))

	fmt.Println(strings.Contains("hellogo","abc"))

	s := []string{"abc","hello","mike","go"}

	buf := strings.Join(s,"-")

	fmt.Println(buf)

	//Index 查找子串的位置
	fmt.Println(strings.Index("abcdhello","hello"))

	fmt.Println(strings.Index("abcdhello","go"))

	buf1 := strings.Repeat("go",2)
	fmt.Println("buf1=",buf1)


	buf3 := "hell@addd@mike"

	fmt.Println(strings.Split(buf3,"@"))

	buf4 := "  are you ok  ?"

	fmt.Println(strings.Fields(buf4))

	for i,data := range strings.Fields(buf4){
		fmt.Println(i,data)
	}
}
