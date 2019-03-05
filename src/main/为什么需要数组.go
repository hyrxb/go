package main

import "fmt"

func main(){
	id1 :=1
	id2 :=2
	id3 :=3

	fmt.Println("id1=%d,id2=%d,id3=%d",id1,id2,id3)

	//同一类型的集合
	var id [50]int

	for i:=0;i<len(id);i++{
		id[i] = i+1
		fmt.Printf("id[%d]=%d\n",i,id[i])
	}

}
