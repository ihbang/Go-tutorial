package main

import (
	"fmt"

	"github.com/ihbang/go-tutorial/bank/accounts"
)

func main() {
	account := accounts.NewAccount("Inhyeok")
	account.Deposit(10)
	fmt.Println(account)

	err := account.Withdraw(15)
	if err != nil {
		fmt.Println(err)
	}
	account.Deposit(10)

	err = account.Withdraw(15)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
