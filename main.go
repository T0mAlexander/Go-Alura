package main

import (
	"fmt"
	"go-bank/accounts"
)

func BillPayment(account checkAcc, billValue float64) {
	account.Withdraw(billValue)
}

type checkAcc interface {
	Withdraw(value float64) string
}

func main() {
	savingAccount := accounts.SavingsAccount{}
	savingAccount.Deposit(100)
	BillPayment(&savingAccount, 60)

	currentAccount := accounts.CurrentAccount{}
	currentAccount.Deposit(500)
	BillPayment(&currentAccount, -400)

	fmt.Println(savingAccount.GetBalance(), currentAccount.GetBalance())
}
