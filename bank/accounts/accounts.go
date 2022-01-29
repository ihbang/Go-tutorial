package accounts

import (
	"errors"
	"fmt"
	"strings"
)

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
	fmt.Print("Deposit: ", amount, " to ", a.Owner(), "'s account\n")
}

//Withdraw amount of money from the account
func (a *account) Withdraw(amount int) error {
	fmt.Print("Try to withdraw ", amount, " from ", a.Owner(), "'s account\n")
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	fmt.Print("Withdraw: ", amount, " from ", a.Owner(), "'s account\n")
	return nil
}

// Balance of the account
func (a account) Balance() int {
	return a.balance
}

// Owner of the account
func (a account) Owner() string {
	return a.owner
}

// String converts account to string
func (a account) String() string {
	strSlice := []string{}
	strSlice = append(strSlice, "--------------------------------------\n")
	strSlice = append(strSlice,
		fmt.Sprint(a.Owner(), "'s account.\n", "Balance: ", a.Balance()))
	strSlice = append(strSlice, "\n--------------------------------------")
	return strings.Join(strSlice, ``)
}
