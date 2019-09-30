package pointers

import (
	"errors"
	"fmt"
)

var (
	ErrInSufficientFunds = errors.New("cannot withdraw, insufficient funds")
)

// Bitcoin type
type Bitcoin int

// String function returns the string value and satisfies the Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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

// Withdraw method withdraws Bitcoin from the wallet and updates balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInSufficientFunds
	}
	w.balance -= amount
	return nil
}
