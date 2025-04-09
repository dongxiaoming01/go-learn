package main

import (
	"fmt"
	"reflect"
)

/*
*
四种基本数据类型：字符串string、正数int、浮点数float64、布尔bool
*/
func main() {
	str := "hello world"
	i := 99
	f := 99.9
	b := true

	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)

	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(b))

}
