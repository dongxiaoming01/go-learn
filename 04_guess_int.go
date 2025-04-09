package main

import (
	"fmt"
	"math/rand/v2"
)

/*
猜测数字游戏：
程序预先设置一个0-10的正整数，用户猜测3次。
*/
func main() {
	target := rand.IntN(10)
	//fmt.Println(target)
	fmt.Println("***猜数字***")
	fmt.Println("规则：三次机会猜测0-10内的数字。")
	var input int
	for i := 0; i < 3; i++ {
		fmt.Printf("请输入数字（第%d次）：\n", i+1)
		fmt.Scanln(&input)
		if input == target {
			fmt.Println("猜对了！")
			return
		} else {
			var tip string
			if input-target > 0 {
				tip = "大"
			} else {
				tip = "小"
			}

			fmt.Printf("猜错了(数字过%s)~~\n", tip)
		}
	}
	fmt.Println("***游戏结束***")
}
