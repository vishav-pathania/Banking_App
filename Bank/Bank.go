package bank

import (
	account "banking_app/Account"
	"banking_app/Error"
	"strings"
)

type Bank struct {
	Bank_id      int
	Fullname     string
	Abbreviation string
	IsActive     bool
	Accounts     []*account.Account
}

func NewBank(bank_id int, Fullname string) (*Bank, *Error.ValidationErr) {
	if bank_id < 0 {
		return nil, Error.NewValidationErr("please provide a valid bank number")
	}
	if Fullname == "" {
		return nil, Error.NewValidationErr("fullname of bank cannot be empty")
	}
	if len(Fullname) < 4 {
		return nil, Error.NewValidationErr("bank fullname cannot be less than 4 letters")
	}
	firstTwo := Fullname[:2]
	lastTwo := Fullname[len(Fullname)-2:]
	Abbreviation := strings.ToLower(firstTwo + lastTwo)
	return &Bank{
		Bank_id:      bank_id,
		Fullname:     Fullname,
		Abbreviation: Abbreviation,
	}, nil

}

func (B *Bank) UpdateBankName(bank_new_name string) *Error.ValidationErr {

	if bank_new_name == "" {
		return Error.NewValidationErr("bank new fullname cannot be empty")
	}
	if len(bank_new_name) < 4 {
		return Error.NewValidationErr("bank new fullname cannot be less than 4 letters")
	}
	B.Fullname = bank_new_name
	firstTwo := bank_new_name[:2]
	lastTwo := bank_new_name[len(bank_new_name)-2:]
	B.Abbreviation = strings.ToLower(firstTwo + lastTwo)
	return nil

}
