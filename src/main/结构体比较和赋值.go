package main

import "fmt"

type Student4 struct{
	id int
	name string
	sex byte
	age int
	addr string
}


func main(){
	s1 := Student4{1,"mike",'m',18,"bj"}
	s2 := Student4{1,"mike",'m',18,"bj"}
	s3 := Student4{2,"mike",'m',18,"bj"}

	fmt.Println("s1 ==s2",s1==s2)
	fmt.Println("s1 ==s3",s1==s3)


	var tmp Student4

	tmp = s3

	fmt.Println("tmp=",tmp)


}
