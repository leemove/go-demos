package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的姓名")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("输入有错误.程序即将推出")
		return
	}
	fmt.Println("Hello " + input + "Welcome!")
}
