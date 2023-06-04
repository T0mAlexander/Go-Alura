package accounts

import "go-bank/customers"

type CurrentAccount struct {
	AccountHolder             customers.AccHolder
	BankBranch, AccountNumber int
	balance                   float64
}

func (Acc *CurrentAccount) BankTransfer(Amount float64, DestinationAcc *CurrentAccount) bool {
	if Amount < Acc.balance && Amount > 0 {
		Acc.balance -= Amount
		DestinationAcc.Deposit(Amount)
		return true
	} else {
		return false
	}
}

func (Acc *CurrentAccount) Withdraw(Amount float64) string {
	canWithdraw := Amount > 0 && Amount <= Acc.balance

	if canWithdraw {
		Acc.balance -= Amount
		return "Sucessful withdrawal!"
	} else {
		return "Insufficient funds!"
	}
}

func (Acc *CurrentAccount) Deposit(Amount float64) (string, float64) {
	if Amount > 0 {
		Acc.balance += Amount
		return "Deposit made sucessfully!", Acc.balance
	} else {
		return "Deposit was denied!", Acc.balance
	}
}

func (acc *CurrentAccount) GetBalance() float64 {
	return acc.balance
}
