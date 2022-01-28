package accounts

import "errors"

var errNoMoney = errors.New("cannot withdraw more than you have")

type account struct {
	owner   string
	balance int
}

// NewAccount creates new account whose owner is name
func NewAccount(name string) *account {
	account := account{owner: name, balance: 0}
	return &account
}

// Deposit amount of money to the account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

//Withdraw amount of money from the account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of the account
func (a account) Balance() int {
	return a.balance
}
