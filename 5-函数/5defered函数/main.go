package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}

func trace(funcName string) func() {
	start := time.Now()
	log.Printf("enter %s", funcName)
	return func() {
		log.Printf("exit %s (%s)", time.Now(), time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
