// main.go
package main

import (
	"fmt"
	"util"
)

func main() {
	util.Welcome()
	MainMenu()
}

func MainMenu() {
MAINFOR:
	for {
		fmt.Println("")
		fmt.Println("*******请选择示例******")
		fmt.Println("1、表示 io.Reader 示例")
		fmt.Println("2、表示 io.ByteReader/ByteWrite 示例")
		fmt.Println("3、退出（q）")
		fmt.Println("**********************")

		var ch string
		fmt.Scanln(&ch)

		switch ch {
		case "1":
			fmt.Println("1")
		case "2":
			fmt.Println("2")
		case "q":
			fmt.Println("程序退出！")
			break MAINFOR
		default:
			fmt.Println("输入错误")
			continue

		}

	}
}
