package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)  // 设定左边树的值
		values = append(values, t.value)       // 把值放入设定范围
		values = appendValues(values, t.right) // 值放入右边的树
	}
	return values
}

func add() {
}

func main() {
	t := new(tree)
	fmt.Print(t)
}
