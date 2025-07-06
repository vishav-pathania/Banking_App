package customer

import (
	account "banking_app/Account"
	bank "banking_app/Bank"
	"banking_app/Error"
	transactions "banking_app/Transactions"
	utils "banking_app/Utils"
	"strconv"
)

type Customer struct {
	Customer_id   int
	First_Name    string
	Last_Name     string
	Accounts      []*account.Account
	Total_Balance float64
}

func NewCustomer(customer_id int, First_Name, Last_Name string, bankaccount *account.Account) (*Customer, *Error.ValidationErr) {
	if customer_id <= 0 {
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

func (C *Customer) AddNewAccount(targetBank *bank.Bank) (*account.Account, *Error.ValidationErr) {
	newAccountId := len(C.Accounts) + 1
	newCustomerAccount, err := account.NewAccount(newAccountId, targetBank)
	if err != nil {
		return nil, err
	}
	C.Accounts = append(C.Accounts, newCustomerAccount)
	C.UpdateTotalBalance()
	return newCustomerAccount, nil
}

func (C *Customer) UpdateTotalBalance() {
	totalSum := 0.0
	for _, CusomterAccountVal := range C.Accounts {
		totalSum += CusomterAccountVal.Balance
	}
	C.Total_Balance = totalSum
}

func (C *Customer) GetAccountById(account_id int) (*account.Account, *Error.ValidationErr) {
	for _, accountVals := range C.Accounts {
		if accountVals.Account_No == account_id {
			return accountVals, nil
		}
	}
	return nil, Error.NewValidationErr("please provide valid account id")
}

func (C *Customer) DeleteAccountById(account_id int) *Error.ValidationErr {
	targetAccount, err := C.GetAccountById(account_id)
	if err != nil {
		return err
	}
	newCustomerAccounts := []*account.Account{}
	for _, AccountVals := range C.Accounts {
		if AccountVals != targetAccount {
			newCustomerAccounts = append(newCustomerAccounts, AccountVals)
		}
	}
	C.Accounts = newCustomerAccounts
	C.UpdateTotalBalance()
	return nil
}

// Two different error type -------
func (C *Customer) DepositMoney(amount float64, account_id int) *Error.TransactionErr {
	targetAccount, err := C.GetAccountById(account_id)
	if err != nil {
		return (*Error.TransactionErr)(err)
	}
	nerr := targetAccount.DepositMoney(amount, C.Customer_id, targetAccount.Account_No)
	if nerr != nil {
		return nerr
	}
	C.UpdateTotalBalance()
	return nil
}

// same doubt as above------
func (C *Customer) WithDrawMoneyByAccount_Id(amount float64, account_id int) *Error.TransactionErr {
	targetAccount, err := C.GetAccountById(account_id)
	if err != nil {
		return (*Error.TransactionErr)(err)
	}
	nerr := targetAccount.WithDrawMoney(amount, C.Customer_id, targetAccount.Account_No)
	if nerr != nil {
		return nerr
	}
	C.UpdateTotalBalance()
	return nil
}

func (C *Customer) GetTotalBalance() float64 {
	TotalBalance := 0.0
	for _, AccountsVal := range C.Accounts {
		TotalBalance += AccountsVal.Balance
	}
	return TotalBalance
}

func (C *Customer) GetAccount_BalanceBy_Id(account_id int) float64 {
	utils.HandlePanic()
	targetAccount, err := C.GetAccountById(account_id)
	if err != nil {
		panic(err)
	}
	return targetAccount.Balance
}

func (C *Customer) TransferMoneyInternally(fromaccount_id, to_account_id int, amount float64) {
	utils.HandlePanic()
	fromAccount, err := C.GetAccountById(fromaccount_id)
	if err != nil {
		panic(err)
	}
	if fromAccount.Balance-amount < 1000 {
		panic("not enough balance to transfer and maintain minimum balance")
	}
	toAccount, err := C.GetAccountById(to_account_id)
	if err != nil {
		panic(err)
	}
	newTransaction, terr := transactions.NewTransaction(amount, C.Customer_id, C.Customer_id, strconv.Itoa(fromAccount.Account_No), strconv.Itoa(toAccount.Account_No))
	if terr != nil {
		panic(terr)
	}
	fromAccount.Transactions = append(fromAccount.Transactions, newTransaction)
	toAccount.Transactions = append(toAccount.Transactions, newTransaction)
	fromAccount.Balance += amount
	toAccount.Balance -= amount
}
