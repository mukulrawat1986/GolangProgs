package pointers

// Bitcoin type
type Bitcoin int

// Wallet type to represent a bitcoin wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit method to deposit bitcoins in the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance method to find the balance in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
