package main

import "fmt"

func main() {
	var i int
	var s string
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&i)
	fmt.Println("请输入一个字符串：")
	fmt.Scanln(&s)
	fmt.Printf("您输入的整数是：%d\n", i)
	fmt.Printf("您输入的字符串是：%s\n", s)
}
