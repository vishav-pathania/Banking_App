package admin

import (
	account "banking_app/Account"
	bank "banking_app/Bank"
	customer "banking_app/Customer"
	transactions "banking_app/Transactions"
	utils "banking_app/Utils"
	"strconv"
)

var adminMap = make(map[int]*Admin)
var adminMapid = -1

type Admin struct {
	Id        int
	FirstName string
	LastName  string
	IsActive  bool
	banks     []*bank.Bank
	customers []*customer.Customer
}

func Newadmin(FirstName, LastName string) *Admin {
	defer utils.HandlePanic()
	if FirstName == "" {
		panic("first name of admin cannot be empty")
	}
	if LastName == "" {
		panic("last name of admin cannot be empty")
	}
	adminMapid++
	ad := &Admin{
		Id:        adminMapid,
		FirstName: FirstName,
		LastName:  LastName,
		IsActive:  true,
		banks:     []*bank.Bank{},
		customers: []*customer.Customer{},
	}
	adminMap[adminMapid] = ad
	return ad
}

func (A *Admin) CreateNewBank(fullname string) *bank.Bank {
	defer utils.HandlePanic()
	newBankId := len(A.banks) + 1
	newBank, err := bank.NewBank(newBankId, fullname)
	if err != nil {
		panic(err)
	}
	return newBank
}

func (A *Admin) UpdateBankName(bank_id int, bank_new_name string) {
	defer utils.HandlePanic()
	targetBank := A.GetBankById(bank_id)
	err := targetBank.UpdateBankName(bank_new_name)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) CreateNewCustomer(First_Name, Last_Name string, Bank_id int) *customer.Customer {

	defer utils.HandlePanic()
	bankobject_to_addcustomer := A.GetBankById(Bank_id)
	newCustomerId := len(A.customers) + 1
	newAccount, err := account.NewAccount(1, bankobject_to_addcustomer)
	if err != nil {
		panic(err)
	}
	newCustomer, err := customer.NewCustomer(newCustomerId, First_Name, Last_Name, newAccount)
	if err != nil {
		panic(err)
	}
	A.customers = append(A.customers, newCustomer)
	return newCustomer
}

func (A *Admin) GetAllCustomers() []customer.Customer {
	allCustomers := []customer.Customer{}
	for _, CustomerVals := range A.customers {
		allCustomers = append(allCustomers, *CustomerVals)
	}
	return allCustomers
}

func (A *Admin) GetBankById(id int) *bank.Bank {
	defer utils.HandlePanic()
	for _, bankVals := range A.banks {
		if bankVals.Bank_id == id {
			return bankVals
		}
	}
	panic("invalid bank id")
}

func (A *Admin) GetAllBanks() []bank.Bank {
	allbanks := []bank.Bank{}
	for _, bankVals := range A.banks {
		allbanks = append(allbanks, *bankVals)
	}
	return allbanks
}

func (A *Admin) GetCustomerById(customer_id int) *customer.Customer {
	defer utils.HandlePanic()
	for _, customerVals := range A.customers {
		if customerVals.Customer_id == customer_id {
			return customerVals
		}
	}
	panic("invalid customer id")
}

func (A *Admin) UpdateCustomer(customer_id int, param string, value interface{}) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	err := targetCustomer.UpdateCustomer(param, value)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) deleteBank(bank_id int) {
	defer utils.HandlePanic()
	targetBank := A.GetBankById(bank_id)
	for _, CustomerVals := range A.customers {
		for _, AccountVals := range CustomerVals.Accounts {
			if AccountVals.Bank == targetBank {
				panic("bank still have accounts associated with it")
			}
		}
	}
	newBanks := []*bank.Bank{}
	for _, banksVal := range A.banks {
		if banksVal != targetBank {
			newBanks = append(newBanks, banksVal)
		}
	}
	A.banks = newBanks
}

func (A *Admin) deleteCustomer(customer_id int) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	if len(targetCustomer.Accounts) > 0 {
		panic("customer still have accounts assosiated with him")
	}
	newCustomers := []*customer.Customer{}
	for _, CustomerVals := range A.customers {
		if CustomerVals != targetCustomer {
			newCustomers = append(newCustomers, CustomerVals)
		}
	}
	A.customers = newCustomers
}

func (A *Admin) AddAccountToCustomer(customer_id int, bank_id int) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	targetBank := A.GetBankById(bank_id)
	_, err := targetCustomer.AddNewAccount(targetBank)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) DeleteCustomerAccountById(customer_id, account_id int) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	err := targetCustomer.DeleteAccountById(account_id)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) DepositMoney(amount float64, customer_id, account_id int) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	err := targetCustomer.DepositMoney(amount, account_id)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) WithDrawMoney(amount float64, customer_id, account_id int) {
	defer utils.HandlePanic()
	targetCustomer := A.GetCustomerById(customer_id)
	err := targetCustomer.WithDrawMoneyByAccount_Id(amount, account_id)
	if err != nil {
		panic(err)
	}
}

func (A *Admin) GetTotalBalanceBy_Customer_Id(customer_id int) float64 {
	targetCustomer := A.GetCustomerById(customer_id)
	TotalBalance := targetCustomer.GetTotalBalance()
	return TotalBalance
}

func (A *Admin) GetAccount_BalanceBy_Id(customer_id, account_id int) float64 {
	targetCustomer := A.GetCustomerById(customer_id)
	accountBalance := targetCustomer.GetAccount_BalanceBy_Id(account_id)
	return accountBalance
}

func (A *Admin) Transfer_MoneyInternally_ByCustomerId(amount float64, customer_id, fromAccount_id, toAccount_id int) {
	targerCustomer := A.GetCustomerById(customer_id)
	targerCustomer.TransferMoneyInternally(fromAccount_id, toAccount_id, amount)
}

func (A *Admin) TransferMoney_To_External(amount float64, fromCustomer_id, ToCustomer_id, fromAccount_id, toAccount_id int) {
	defer utils.HandlePanic()
	senderCustomer := A.GetCustomerById(fromCustomer_id)
	receiverCustomer := A.GetCustomerById(ToCustomer_id)
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
}
