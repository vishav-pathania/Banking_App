package transactions

import (
	"banking_app/Error"
	utils "banking_app/Utils"
)

type Transaction struct {
	Transaction_Id          int64
	Amount                  float64
	FromCustomer_id         int
	ToCustomer_id           int
	FromCustomer_Account_id string
	ToCustomer_Account_id   string
}

func NewTransaction(amount float64, FromCustomer_id, ToCustomer_id int, FromCustomer_Account_id string, ToCustomer_Account_id string) (*Transaction, *Error.TransactionErr) {
	if amount <= 0 {
		return nil, Error.NewTransactionErr("transaction amount should be greater than zero")
	}
	if FromCustomer_Account_id == ToCustomer_Account_id {
		return nil, Error.NewTransactionErr("from_account and to_account cannot be similar")
	}
	return &Transaction{
		Transaction_Id:          utils.GenerateTransactionID(),
		Amount:                  amount,
		FromCustomer_id:         FromCustomer_id,
		ToCustomer_id:           ToCustomer_id,
		FromCustomer_Account_id: FromCustomer_Account_id,
		ToCustomer_Account_id:   ToCustomer_Account_id,
	}, nil

}
