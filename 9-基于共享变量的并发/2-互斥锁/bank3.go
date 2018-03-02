package bank3

import (
	"sync"
)

// 只有一个容量的chan会保证deposit 和balance只有一个在运行
var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
