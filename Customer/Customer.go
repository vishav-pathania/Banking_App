package customer

import (
	account "banking_app/Account"
	"banking_app/Error"
	utils "banking_app/Utils"
)

type Customer struct {
	Customer_id   int
	First_Name    string
	Last_Name     string
	Accounts      []*account.Account
	Total_Balance int
}

func NewCustomer(customer_id int, First_Name, Last_Name string, bankaccount *account.Account) (*Customer, *Error.ValidationErr) {
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

func (C *Customer) UpdateCustomer(param string, value interface{}) *Error.ValidationErr {
	switch param {
	case "First_Name":
		err := C.UpdateFirstName(value)
		if err != nil {
			return err
		}
		return nil
	case "Last_Name":
		err := C.UpdateLastName(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return Error.NewValidationErr("no matching params found to update")
	}
}

func (C *Customer) UpdateFirstName(value interface{}) *Error.ValidationErr {
	if utils.GetVariableType(value) != "string" {
		return Error.NewValidationErr("please enter a string value")
	}
	if value == "" {
		return Error.NewValidationErr("first name cannot be set to empty")
	}
	conval, ok := value.(string)
	if !ok {
		return Error.NewValidationErr("error in setting First_Name string")
	}
	C.First_Name = conval
	return nil
}

func (C *Customer) UpdateLastName(value interface{}) *Error.ValidationErr {
	if utils.GetVariableType(value) != "string" {
		return Error.NewValidationErr("please enter a string value")
	}
	if value == "" {
		return Error.NewValidationErr("last name cannot be set to empty")
	}
	conval, ok := value.(string)
	if !ok {
		return Error.NewValidationErr("error in setting Last_Name string")
	}
	C.First_Name = conval
	return nil
}
