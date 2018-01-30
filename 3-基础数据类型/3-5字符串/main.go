package main

import (
	"fmt"
	"strings"
)

func main() {
	str := basename("a/d/adsf/txt.1234")
	fmt.Printf(str)
}

// func basename(str string) string {
// 	// 删除,路径部分
// 	for i := len(str) - 1; i >= 0; i-- {
// 		if str[i] == '/' {
// 			str = str[i+1:]
// 			break
// 		}
// 	}
// 	// 删除后缀
// 	for i := len(str) - 1; i >= 0; i-- {
// 		if str[i] == '.' {
// 			str = str[:i]
// 		}
// 	}
// 	return str
// }

func basename(str string) string {
	slash := strings.LastIndex(str, "/")
	str = str[slash+1:]

	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}

	return str
}
