package main

import (
	"fmt"
	"my-go-app/example/accounts"
	// "my-go-app/example/hello"
	// "myFunc.com/ufunc"
)

func main() {
	// hello.PrintHello()
	// ufunc.PrintFunc()

	account := accounts.NewAccount("nicolas")
	account.Deposit(10)
	println(*account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	println(*account.Balance())
}
