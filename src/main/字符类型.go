package main

import "fmt"

func main() {
	var ch byte //声明字符类型
	ch = 97

	fmt.Println("ch =", ch)

	//格式话输出　,%c以字符串方式打印，%d以整数方式打印

	ch = 'a' //字符，单引号

	fmt.Printf("%c,%d\n", ch, ch)

	//大写转小写，小写转大写，大小写相差32 大写:65,小写:97
	fmt.Printf("大写:%d,小写:%d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A'+32)
	fmt.Printf("小写转大写:%c\n", 'a'-32)

	//以 ‘\’ 反斜杠　开头的是转义字符

}
