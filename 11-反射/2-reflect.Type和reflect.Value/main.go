package main

import (
	"fmt"
	"reflect"
)

func main() {
	// t := reflect.TypeOf(3)
	// fmt.Println(t)
	// fmt.Println(t.String())
	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Print(i)
}
