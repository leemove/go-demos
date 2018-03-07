package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	readFile("document.txt")
}

func readFile(url string) {
	inputFile, err := os.Open(url)
	if err != nil {
		fmt.Println("读取文件失败!")
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	var index int
	for {
		index++
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("第%d行:%s", index, inputString)
		if err == io.EOF {
			return
		}
	}
}
