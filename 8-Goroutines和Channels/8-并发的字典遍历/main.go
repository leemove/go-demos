package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("p", true, "查看进度")
var sema = make(chan struct{}, 20)

func main() {
	var tick <-chan time.Time
	var n sync.WaitGroup
	// timeStart := time.Now()
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func(roots []string) {
		for _, root := range roots {
			n.Add(1)
			go walkDir(root, fileSizes, &n)
		}
	}(roots)

	go func() {
		n.Wait()
		close(fileSizes)
		// span := time.Since(timeStart)
		// fmt.Println(span)
	}()

	var nfiles, nbytes int64
	// for size := range fileSizes {
	// 	nfiles++
	// 	nbytes += size
	// }
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d个文件, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
