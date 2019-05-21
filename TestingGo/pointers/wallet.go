package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin type based on int
type Bitcoin int

// String method returns the representation of Bitcoin, thus implementing
// the Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet type to hold our bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit method deposits bitcoins in our wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance method returns the number of bitcoins in our wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw method to withdraw bitcoins from our wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
