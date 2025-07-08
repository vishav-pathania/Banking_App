package ledger

import "banking_app/Error"

type Ledger struct {
	From_bank_id   int
	From_bank_name string
	To_bank_id     int
	To_bank_name   string
	Amount         float64
}

func Newledger(amount float64, from_bank_id, to_bank_id int, from_bank_name, to_bank_name string) (*Ledger, *Error.TransactionErr) {
	if from_bank_id == to_bank_id {
		return nil, Error.NewTransactionErr("from_bank and to_bank cannot be same")
	}
	return &Ledger{
		From_bank_id:   from_bank_id,
		From_bank_name: from_bank_name,
		To_bank_id:     to_bank_id,
		To_bank_name:   to_bank_name,
		Amount:         amount,
	}, nil
}
