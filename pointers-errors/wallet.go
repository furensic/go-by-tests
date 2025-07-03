package pointers_errors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin uint

type Wallet struct {
	balance Bitcoin
}

var (
	ErrorInsufficientFunds = errors.New("asd")
)

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
