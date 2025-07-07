package account

import (
	"banking_app/Error"
	transactions "banking_app/Transactions"
	"strconv"
)

type Account struct {
	Account_No   int
	Balance      float64
	Bank_id      int
	Transactions []*transactions.Transaction
	IsActive     bool
}

func NewAccount(Account_No int, Bank_id int) (*Account, *Error.ValidationErr) {

	return &Account{
		Account_No: Account_No,
		Balance:    1000,
		Bank_id:    Bank_id,
		IsActive:   true,
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
	if A.Balance-amount < 1000 {
		return Error.NewTransactionErr("must maintain a minimum balance of 1000/-")
	}
	newTransaction, err := transactions.NewTransaction(amount, FromCustomer_id, FromCustomer_id, strconv.Itoa(FromCustomer_AccountNo), "WithDraw")
	if err != nil {
		return err
	}
	A.Transactions = append(A.Transactions, newTransaction)
	A.Balance -= amount
	return nil
}
