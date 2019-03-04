package lesson1

import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	fmt.Println(Books{"Go 语言","go.zxb.cc","自学go",1222})
	fmt.Println(Books{"python语言","python.zxb8.cc","自学python",1223})
	fmt.Println(Books{"php语言","php.zxb8.cc","自学php",1111})
}
