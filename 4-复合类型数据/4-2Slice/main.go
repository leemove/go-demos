package main

import (
	"fmt"
)

var runes []rune

func main() {
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q \n", runes)
}
