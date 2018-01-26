package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	for _, filename := range files {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %e /n", err)
		} else {
			stringData := string(data)
			for _, word := range strings.Split(stringData, "\n") {
				fmt.Println(word)
			}
		}
	}
}
