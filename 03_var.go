package main

import "fmt"

/*
字面量literal、变量variable、常量constant

声明 var s string
赋值 s = "hello"
*/
func main() {
	fmt.Println("字面量")
	var a string = "声明+赋值"
	fmt.Println(a)
	b := "短声明，依据值确定类型"
	fmt.Println(b)
	const c = "常量"
	fmt.Println(c)
}
