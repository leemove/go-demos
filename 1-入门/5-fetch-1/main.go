package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		if strings.HasPrefix(url, "http://") {
			fetch(url)
		} else {
			url = "http://" + url
			fetch(url)
		}
	}
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "获取网址失败!地址:%s,错误:%v", url, err)
	}
	nbytes, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取数据失败!地址:%s,错误:%v", url, err)
	}
	fmt.Printf("数据长度为:%7d", nbytes)
}
