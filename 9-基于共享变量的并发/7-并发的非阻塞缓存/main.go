package main

import (
	"fmt"
	"godemos/9-基于共享变量的并发/7-并发的非阻塞缓存/memo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://gopl-zh.b0.upaiyun.com/ch9/ch9-07.html", "http://blog.csdn.net/chenbaoke/article/details/41957725", "http://blog.csdn.net/chenbaoke/article/details/41957725", "http://blog.csdn.net/chenbaoke/article/details/41957725"}
	m := memo.New(httpGetBody)
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
