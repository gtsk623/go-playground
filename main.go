package main

import (
	"fmt"
	"log"

	"github.com/gtsk623/go-playground/accounts"
)

func main() {
	account := accounts.NewAccount("test")
	fmt.Println(account.Balance())
	res := account.Withdraw(5)
	if res != nil {
		fmt.Println(res)
	}
	account.Deposit(20)
	fmt.Println(account.Balance())
	res2 := account.Withdraw(10)
	if res2 != nil {
		fmt.Println(res2)
		log.Fatalln(res2)
	}
	fmt.Println(account.Balance())
}
