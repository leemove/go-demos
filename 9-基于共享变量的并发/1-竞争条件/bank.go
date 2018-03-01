package bank

var deposits = make(chan int) // 存款
var balances = make(chan int) //

func Deposit(amount int) {
	deposits <- amount
}

func Blance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
