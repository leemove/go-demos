package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str1 = "asSASA ddd dsjkdsjs dk"
	const str2 = "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("字节数分别为: %d, %d\n", len(str1), len(str2))                                       // 22 28 说明这两个日文长度为(28-22)/2 = 3
	fmt.Printf("字符数分别为: %d, %d\n", utf8.RuneCountInString(str1), utf8.RuneCountInString(str2)) // 22 24 真正的字符数
}
