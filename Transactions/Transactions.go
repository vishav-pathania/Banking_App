package transactions

import (
	"banking_app/Error"
	utils "banking_app/Utils"
)

type Transaction struct {
	Transaction_Id       int64
	Amount               int
	FromCustomer_id      int
	ToCustomer_id        int
	FromCustomer_Account string
	ToCustomer_Account   string
}

func NewTransaction(amount, FromCustomer_id, ToCustomer_id int, FromCustomer_Account, ToCustomer_Account string) (*Transaction, *Error.TransactionErr) {
	if amount <= 0 {
		return nil, Error.NewTransactionErr("transaction amount should be greater than zero")
	}
	if FromCustomer_Account == ToCustomer_Account {
		return nil, Error.NewTransactionErr("from_account and to_account cannot be similar")
	}
	return &Transaction{
		Transaction_Id:       utils.GenerateTransactionID(),
		Amount:               amount,
		FromCustomer_id:      FromCustomer_id,
		ToCustomer_id:        ToCustomer_id,
		FromCustomer_Account: FromCustomer_Account,
		ToCustomer_Account:   ToCustomer_Account,
	}, nil

}
