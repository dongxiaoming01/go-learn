package main

import (
	"fmt"
	"time"
)

func work(i int) {
	fmt.Printf("start:%d\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("end:%d\n", i)
}
func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}

/*
go关键字，将使用协程运行方法；
defer关键字，等待函数完成前才执行，多个defer会反向顺序执行。
*/
func normal_or_go(use_go bool) {
	defer timeCost()()
	if use_go {
		for i := 0; i < 5; i++ {
			go work(i)
		}
	} else {
		for i := 0; i < 5; i++ {
			work(i)
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("main thread\n")
}

func main() {
	normal_or_go(false)
	normal_or_go(true)
}
