package main

import (
	"errors"
	"sync"
)

type Account struct {
	Balance int64
	Id      string
	mu      sync.Mutex
}

func (account *Account) Transfer(from, to *Account, amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if from == nil || to == nil {
		return errors.New("from and to must both be specified")
	}
	if from.Id == to.Id {
		return errors.New("from and to id cannot be the same account")
	}
	first, second := from, to
	if from.Id < to.Id {
		first, second = second, from
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	if from.Balance < amount {
		return errors.New("not enough funds")
	}

	from.Balance -= amount
	to.Balance += amount

	return nil
}

func main() {

}
