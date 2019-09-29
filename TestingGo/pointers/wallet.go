package pointers

// Wallet type to represent a bitcoin wallet
type Wallet struct {
	balance int
}

// Deposit method to deposit bitcoins in the wallet
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance method to find the balance in the wallet
func (w Wallet) Balance() int {
	return w.balance
}
