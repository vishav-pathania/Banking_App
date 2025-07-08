package customer

import (
	account "banking_app/Account"
	bank "banking_app/Bank"
	"banking_app/Error"
	ledger "banking_app/Ledger"
	transactions "banking_app/Transactions"
	utils "banking_app/Utils"
	"strconv"
)

var customerMap = make(map[int]*Customer)
var bankMap = make(map[int]*bank.Bank)

type Customer struct {
	Customer_id   int
	First_Name    string
	Last_Name     string
	Accounts      []*account.Account
	Total_Balance float64
	isAdmin       bool
	isActive      bool
}

func newCustomer(First_Name, Last_Name string, isAdmin bool, bankaccount *account.Account) (*Customer, *Error.ValidationErr) {

	if First_Name == "" {
		return nil, Error.NewValidationErr("first name of customer cannot be empty")
	}
	if Last_Name == "" {
		return nil, Error.NewValidationErr("last name of customer cannot be empty")
	}
	newCustomerId := utils.GenerateUniqueID()
	if !isAdmin {
		newTransaction, terr := transactions.NewTransaction(1000, newCustomerId, newCustomerId, strconv.Itoa(bankaccount.Account_No), "New-Account-Deposite")
		if terr != nil {
			return nil, (*Error.ValidationErr)(terr)
		}
		bankaccount.Transactions = append(bankaccount.Transactions, newTransaction)
	}

	return &Customer{
		Customer_id:   newCustomerId,
		First_Name:    First_Name,
		Last_Name:     Last_Name,
		Accounts:      []*account.Account{bankaccount},
		Total_Balance: 1000,
		isAdmin:       isAdmin,
		isActive:      true,
	}, nil
}

func Newadmin(FirstName, LastName string) *Customer {
	defer utils.HandlePanic()

	newadmin, err := newCustomer(FirstName, LastName, true, nil)
	if err != nil {
		panic(err)
	}
	return newadmin
}

func (C *Customer) CreateNewBank(fullname string) *bank.Bank {
	defer utils.HandlePanic()
	if !C.isActive {
		panic("admin needs to be active to create a bank")
	}
	if !C.isAdmin {
		panic("only admin can create a bank")
	}
	newBankId := utils.GenerateUniqueID()
	newBank, err := bank.NewBank(newBankId, fullname)
	if err != nil {
		panic(err)
	}

	bankMap[newBankId] = newBank
	return newBank
}

func (C *Customer) UpdateBankName(bank_id int, bank_new_name string) {

	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can create a bank")
	}
	if !C.isActive {
		panic("admin needs to be active to update a bank")
	}
	targetBank := C.GetBankById(bank_id)
	err := targetBank.UpdateBankName(bank_new_name)
	if err != nil {
		panic(err)
	}
}

func (C *Customer) CreateNewCustomer(First_Name, Last_Name string, Bank_id int) *Customer {

	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can create a new customer")
	}
	if !C.isActive {
		panic("only active admins can create a new customer")
	}
	targetBank := C.GetBankById(Bank_id)
	newAccount, err := account.NewAccount(utils.GenerateUniqueID(), Bank_id)
	if err != nil {
		panic(err)
	}

	targetBank.Accounts = append(targetBank.Accounts, newAccount)
	newCustomer, err := newCustomer(First_Name, Last_Name, false, newAccount)
	if err != nil {
		panic(err)
	}
	customerMap[newCustomer.Customer_id] = newCustomer
	return newCustomer
}

func (C *Customer) GetAllCustomers() []Customer {
	utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can get all customers")
	}
	if !C.isActive {
		panic("only active admins can get all customers")
	}
	allCustomers := []Customer{}
	for _, CustomerVals := range customerMap {
		allCustomers = append(allCustomers, *CustomerVals)
	}
	return allCustomers
}

func (C *Customer) GetBankById(id int) *bank.Bank {
	defer utils.HandlePanic()
	targetBank, ok := bankMap[id]
	if !ok {
		panic("invalid bank id")
	}
	return targetBank
}

func (C *Customer) GetAllBanks() []bank.Bank {
	if !C.isAdmin {
		panic("only admins can get all banks")
	}
	if !C.isActive {
		panic("only active admins can get all banks")
	}
	allbanks := []bank.Bank{}
	for _, bankVals := range bankMap {
		allbanks = append(allbanks, *bankVals)
	}
	return allbanks
}

func (C *Customer) GetCustomerById(customer_id int) *Customer {
	defer utils.HandlePanic()
	targetCustomer, ok := customerMap[customer_id]
	if !ok {
		panic("invalid customer id")
	}
	if !targetCustomer.isActive {
		panic("invalid customer id")
	}
	return targetCustomer
}

func (C *Customer) DeleteBank(bank_id int) {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admins can delete a account")
	}
	if !C.isActive {
		panic("only active admins can delete a account")
	}
	targetBank := C.GetBankById(bank_id)
	for _, accountVals := range targetBank.Accounts {
		if accountVals.IsActive {
			panic("banks with active accounts cannot be deleted")
		}
	}
	targetBank.IsActive = false
}

func (C *Customer) DeleteCustomer(customer_id int) {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can delete a customer")
	}
	if !C.isActive {
		panic("only active admin can delete a customer")
	}
	targetCustomer := C.GetCustomerById(customer_id)
	if len(targetCustomer.Accounts) > 0 {
		panic("customer still have accounts assosiated with him")
	}
	targetCustomer.isActive = false
}

func (C *Customer) DeleteCustomerAccountById(customer_id, account_id int) {
	defer utils.HandlePanic()
	targetCustomer := C.GetCustomerById(customer_id)
	err := targetCustomer.DeleteAccountById(account_id)
	if err != nil {
		panic(err)
	}
}

// func (C *Customer) DepositMoney(amount float64, customer_id, account_id int) {
// 	defer utils.HandlePanic()
// 	targetCustomer := C.GetCustomerById(customer_id)
// 	err := targetCustomer.DepositMoney(amount, account_id)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (C *Customer) WithDrawMoney(amount float64, account_id int) {
	defer utils.HandlePanic()
	err := C.WithDrawMoneyByAccount_Id(amount, account_id)
	if err != nil {
		panic(err)
	}
}

func (C *Customer) GetTotalBalanceBy_Customer_Id(customer_id int) float64 {
	targetCustomer := C.GetCustomerById(customer_id)
	TotalBalance := targetCustomer.GetTotalBalance()
	return TotalBalance
}

// func (C *Customer) GetAccount_BalanceBy_Id(customer_id, account_id int) float64 {
// 	targetCustomer := C.GetCustomerById(customer_id)
// 	accountBalance := targetCustomer.GetAccount_BalanceBy_Id(account_id)
// 	return accountBalance
// }

// func (C *Customer) Transfer_MoneyInternally_ByCustomerId(amount float64, customer_id, fromAccount_id, toAccount_id int) {
// 	targerCustomer := C.GetCustomerById(customer_id)
// 	targerCustomer.TransferMoneyInternally(fromAccount_id, toAccount_id, amount)
// }

func (C *Customer) TransferMoney_To_External(amount float64, fromCustomer_id, ToCustomer_id, fromAccount_id, toAccount_id int) {
	defer utils.HandlePanic()
	senderCustomer := C.GetCustomerById(fromCustomer_id)
	receiverCustomer := C.GetCustomerById(ToCustomer_id)
	senderAccount, Aserr := senderCustomer.GetAccountById(fromAccount_id)
	if Aserr != nil {
		panic(Aserr)
	}
	receiverAccount, Arerr := receiverCustomer.GetAccountById(toAccount_id)
	if Arerr != nil {
		panic(Arerr)
	}
	if senderAccount.Balance-amount < 1000 {
		panic("not enough balance to transfer and maintain minimum balance")
	}
	newTransaction, terr := transactions.NewTransaction(amount, senderCustomer.Customer_id, receiverCustomer.Customer_id, strconv.Itoa(senderAccount.Account_No), strconv.Itoa(receiverAccount.Account_No))
	if terr != nil {
		panic(terr)
	}
	senderAccount.Transactions = append(senderAccount.Transactions, newTransaction)
	receiverAccount.Transactions = append(receiverAccount.Transactions, newTransaction)
	senderAccount.Balance -= amount
	receiverAccount.Balance += amount

	senderCustomer.UpdateTotalBalance()
	receiverCustomer.UpdateTotalBalance()
	/////////////////////////////////// adding ledger below
	MoneySendingBank := C.GetBankById(senderAccount.Bank_id)
	MoneyReceivingBank := C.GetBankById(receiverAccount.Bank_id)
	newEntryInLedger, err := ledger.Newledger(amount, MoneySendingBank.Bank_id, MoneyReceivingBank.Bank_id, MoneySendingBank.Fullname, MoneyReceivingBank.Fullname)
	if err != nil && MoneySendingBank.Bank_id != MoneyReceivingBank.Bank_id {
		panic(err)
	}
	MoneySendingBank.Ledger = append(MoneySendingBank.Ledger, newEntryInLedger)
	MoneyReceivingBank.Ledger = append(MoneyReceivingBank.Ledger, newEntryInLedger)
}

func (C *Customer) GetPassBook_ById(customer_id, account_id int, pageNo int) []transactions.Transaction {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admins can get specific passbooks by id")
	}
	if !C.isActive {
		panic("only active admins can get specific passbooks by id")
	}
	defer utils.HandlePanic()
	targetCustomer := C.GetCustomerById(customer_id)
	targetAccount, err := targetCustomer.GetAccountById(account_id)
	if err != nil {
		panic(err)
	}
	pageSize := 5
	startIndex := pageNo * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(targetAccount.Transactions) {
		endIndex = len(targetAccount.Transactions)
	}
	if startIndex >= len(targetAccount.Transactions) {
		panic("provide a smaller page number transactions doesn't exist")
	}
	copyOfTransactions := []transactions.Transaction{}
	for i := startIndex; i < endIndex; i++ {
		copyOfTransactions = append(copyOfTransactions, *targetAccount.Transactions[i])
	}
	return copyOfTransactions
}

func (C *Customer) UpdateCustomer(param string, value interface{}) *Error.ValidationErr {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can update customers")
	}
	if !C.isActive {
		panic("only active admins can update customers")
	}
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

func (C *Customer) AddNewAccount(bank_id int) *account.Account {
	defer utils.HandlePanic()
	newAccountId := utils.GenerateUniqueID()
	targetBank := C.GetBankById(bank_id)
	newCustomerAccount, err := account.NewAccount(newAccountId, bank_id)
	if err != nil {
		panic(err)
	}
	C.Accounts = append(C.Accounts, newCustomerAccount)
	targetBank.Accounts = append(targetBank.Accounts, newCustomerAccount)
	newTransaction, terr := transactions.NewTransaction(1000, C.Customer_id, C.Customer_id, strconv.Itoa(newCustomerAccount.Account_No), "New-Account-Deposite")
	if terr != nil {
		panic(err)
	}
	newCustomerAccount.Transactions = append(newCustomerAccount.Transactions, newTransaction)
	C.UpdateTotalBalance()
	return newCustomerAccount
}

func (C *Customer) UpdateTotalBalance() {
	totalSum := 0.0
	for _, CusomterAccountVal := range C.Accounts {
		if CusomterAccountVal.IsActive {
			totalSum += CusomterAccountVal.Balance
		}
	}
	C.Total_Balance = totalSum
}

func (C *Customer) GetAccountById(account_id int) (*account.Account, *Error.ValidationErr) {
	for _, accountVals := range C.Accounts {
		if accountVals.Account_No == account_id && accountVals.IsActive {
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
	targetAccount.IsActive = false
	C.UpdateTotalBalance()
	return nil
}

func (C *Customer) DepositMoney(amount float64, account_id int) {
	defer utils.HandlePanic()
	/////////////////////////////
	if !C.isActive {
		panic("inactive customers cannot deposit money")
	}
	targetAccount, err := C.GetAccountById(account_id)
	if err != nil {
		panic(err)
	}
	nerr := targetAccount.DepositMoney(amount, C.Customer_id, targetAccount.Account_No)
	if nerr != nil {
		panic(nerr)
	}
	C.UpdateTotalBalance()

}

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
	if !C.isActive {
		panic("inactive customers cannot transfer money")
	}
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
	if !fromAccount.IsActive || !toAccount.IsActive {
		panic("sending account cannot be inactive")
	}
	if !toAccount.IsActive {
		panic("receiving account cannot be inactive")
	}
	newTransaction, terr := transactions.NewTransaction(amount, C.Customer_id, C.Customer_id, strconv.Itoa(fromAccount.Account_No), strconv.Itoa(toAccount.Account_No))
	if terr != nil {
		panic(terr)
	}
	fromAccount.Transactions = append(fromAccount.Transactions, newTransaction)
	toAccount.Transactions = append(toAccount.Transactions, newTransaction)
	fromAccount.Balance -= amount
	toAccount.Balance += amount
}

func (C *Customer) GetLedgerByBank_Id(bank_id, pageNo int) []ledger.Ledger {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admins can get specific bank ledger by id")
	}
	if !C.isActive {
		panic("only active admins can get specific bank ledger by id")
	}
	targetBank := C.GetBankById(bank_id)
	pageSize := 5
	startIndex := pageNo * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(targetBank.Ledger) {
		endIndex = len(targetBank.Ledger)
	}
	if startIndex >= len(targetBank.Ledger) {
		panic("provide a smaller page number ledger entry doesn't exist")
	}
	copyOfBankLedger := []ledger.Ledger{}
	for i := startIndex; i < endIndex; i++ {
		copyOfBankLedger = append(copyOfBankLedger, *targetBank.Ledger[i])
	}
	return copyOfBankLedger
}

func (C *Customer) SettleMent(fromBank_id, toBank_id int) float64 {
	defer utils.HandlePanic()
	if !C.isAdmin {
		panic("only admin can get bank settlements")
	}
	if !C.isActive {
		panic("only active admins can get bank settlements")
	}
	fromBank := C.GetBankById(fromBank_id)
	_ = C.GetBankById(toBank_id) //just for validation
	sendingAmount := 0.0
	receivingAmount := 0.0
	for _, LedgerVals := range fromBank.Ledger {
		if LedgerVals.From_bank_id == fromBank_id {
			sendingAmount += LedgerVals.Amount
		}
		if LedgerVals.To_bank_id == fromBank_id {
			receivingAmount += LedgerVals.Amount
		}
	}
	return sendingAmount - receivingAmount
}
