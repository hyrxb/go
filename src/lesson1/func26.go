package lesson1

import "fmt"

type Books1 struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	var book1 Books1
	var book2 Books1

	book1.title = "GO语言"
	book1.author = "go.zxb8.cc"
	book1.subject= "自学go"
	book1.book_id =1

	book2.title = "php语言"
	book2.author="php.zxb8.cc"
	book2.subject = "自学php"
	book2.book_id = 2

	fmt.Printf("book1 title:%s\n",book1.title)
	fmt.Printf("book1 author:%s\n",book1.author)
	fmt.Printf("book1 subject:%s\n",book1.subject)
	fmt.Printf("book1 book_id{%d\n",book1.book_id)

	fmt.Printf("book2 title:%s\n",book2.title)
	fmt.Printf("book2 author:%s\n",book2.title)
	fmt.Printf("book2 subject:%s\n",book2.subject)
	fmt.Printf("book2 book_id:%d\n",book2.book_id)
}
