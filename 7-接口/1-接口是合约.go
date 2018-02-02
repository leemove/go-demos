package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	name := "Leemove"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Print(c)
}
