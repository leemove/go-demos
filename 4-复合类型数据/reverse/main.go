package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	reverse(&arr)
	fmt.Printf("%d", arr)
}

func reverse(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}
