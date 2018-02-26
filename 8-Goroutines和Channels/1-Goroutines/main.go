package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	// 在main函数退出后 所有的gorun
	const n = 45
	fibN := fib(n)
	fmt.Printf("\r斐波那契数列第%d项值为%d \n", n, fibN)
}

func fib(x int) int {
	// 递归函数比较慢
	if x < 2 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
