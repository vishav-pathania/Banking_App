package customer

import (
	account "banking_app/Account"
	"banking_app/Error"
)

type Customer struct {
	Customer_id   float64
	First_Name    string
	Last_Name     string
	Accounts      []*account.Account
	Total_Balance float64
}

func NewCustomer(customer_id float64, First_Name, Last_Name string, bankaccount *account.Account) (*Customer, *Error.ValidationErr) {
	if customer_id <= 999999999 {
		return nil, Error.NewValidationErr("please provide a valid customer_id number")
	}
	if First_Name == "" {
		return nil, Error.NewValidationErr("first name of customer cannot be empty")
	}
	if Last_Name == "" {
		return nil, Error.NewValidationErr("last name of customer cannot be empty")
	}
	return &Customer{
		Customer_id:   customer_id,
		First_Name:    First_Name,
		Last_Name:     Last_Name,
		Accounts:      []*account.Account{bankaccount},
		Total_Balance: 1000,
	}, nil
}
