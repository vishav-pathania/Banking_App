package bank

import (
	"banking_app/Error"
	"strings"
)

type Bank struct {
	Bank_id      int
	Fullname     string
	Abbreviation string
}

func NewBank(bank_id int, Fullname string) (*Bank, *Error.ValidationErr) {
	if bank_id <= 9999 {
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
