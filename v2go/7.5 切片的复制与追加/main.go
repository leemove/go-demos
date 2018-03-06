package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	arr = Filter(arr, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println()
}

func Filter(s []int, fn func(int) bool) []int {
	var newArr []int
	for _, number := range s {
		if fn(number) {
			newArr = append(newArr, number)
		}
	}
	return newArr
}
