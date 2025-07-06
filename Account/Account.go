package account

import (
	bank "banking_app/Bank"
	"banking_app/Error"
	transactions "banking_app/Transactions"
)

type Account struct {
	Account_No int
	*bank.Bank
	Balance      float64
	Transactions []*transactions.Transaction
}

func NewAccount(Account_No int, bankobject *bank.Bank) (*Account, *Error.ValidationErr) {
	if Account_No <= 0 {
		return nil, Error.NewValidationErr("please provide a valid account number")
	}
	return &Account{
		Account_No: Account_No,
		Bank:       bankobject,
		Balance:    1000,
	}, nil
}
