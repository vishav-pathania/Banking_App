package account

import (
	bank "banking_app/Bank"
	"banking_app/Error"
	transactions "banking_app/Transactions"
	"strconv"
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

func (A *Account) DepositMoney(amount float64, FromCustomer_id int, ToCustomer_AccountNo int) *Error.TransactionErr {
	newTransaction, err := transactions.NewTransaction(amount, FromCustomer_id, FromCustomer_id, "Self-Deposit", strconv.Itoa(ToCustomer_AccountNo))
	if err != nil {
		return err
	}
	A.Transactions = append(A.Transactions, newTransaction)
	A.Balance += amount
	return nil
}

func (A *Account) WithDrawMoney(amount float64, FromCustomer_id int, FromCustomer_AccountNo int) *Error.TransactionErr {
	newTransaction, err := transactions.NewTransaction(amount, FromCustomer_id, FromCustomer_id, strconv.Itoa(FromCustomer_AccountNo), "WithDraw")
	if err != nil {
		return err
	}
	A.Transactions = append(A.Transactions, newTransaction)
	A.Balance -= amount
	return nil
}
