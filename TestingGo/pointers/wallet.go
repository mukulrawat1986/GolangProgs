package wallet

// Wallet type to hold our bitcoins
type Wallet struct {
	balance int
}

// Deposit method deposits bitcoins in our bitcoin
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance method returns the number of bitcoins in our wallet
func (w *Wallet) Balance() int {
	return w.balance
}
