package wallet

// Bitcoin type based on int
type Bitcoin int

// Wallet type to hold our bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit method deposits bitcoins in our bitcoin
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance method returns the number of bitcoins in our wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
