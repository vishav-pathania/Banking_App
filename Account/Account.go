package account

import (
	bank "banking_app/Bank"
	"banking_app/Error"
)

type Account struct {
	Account_No float64
	bank.Bank
	Balance float64
}

func NewAccount(Account_No float64, bankobject bank.Bank) (*Account, *Error.ValidationErr) {
	if Account_No <= 999999999 {
		return nil, Error.NewValidationErr("please provide a valid account number")
	}
	return &Account{
		Account_No: Account_No,
		Bank:       bankobject,
		Balance:    1000,
	}, nil
}
