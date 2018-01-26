package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	for _, fileName := range files {
		data, err := ioutil.ReadFile(fileName)
		kmap := make(map[string]int)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
		} else {
			stringdata := string(data)
			for _, lineStrings := range strings.Split(stringdata, "\n") {
				kmap[lineStrings]++
			}
		}
		for key, value := range kmap {
			// fmt.Printf("key:%s,value:%d\n", key, value)
			if value > 1 {
				fmt.Printf("in file %s,string: %s has been fount %d times\n", fileName, key, value)
			}
		}
	}
}
