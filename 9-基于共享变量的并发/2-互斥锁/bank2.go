package bank2

// 只有一个容量的chan会保证deposit 和balance只有一个在运行
var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
