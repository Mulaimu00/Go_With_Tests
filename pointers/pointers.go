package pointers

import (
	"errors"
	"fmt"
)


type Bitcoin int
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(depoAmount Bitcoin)  {
	w.balance += depoAmount
}

func (w *Wallet) Withdraw(withdrawAmount Bitcoin) error{
	if withdrawAmount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= withdrawAmount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
