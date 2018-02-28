// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+
func main() {
	// fmt.Println("Commencing countdown.")
	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }
	// launch()
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("开始倒计时,enter键取消倒计时")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("倒计时被取消")
		return
	}
	launch()
}

//!-

func launch() {
	fmt.Println("开炮!!!")
}
