package main

import (
	"fmt"
)

type IntSet int

// 挂载在指针上和挂载在对象上的方法是有区别的.
func (i *IntSet) String() string {
	return "123123"
}

var x IntSet

// var _ = x.String()

var _ fmt.Stringer = &x
