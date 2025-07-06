package main

import (
	admin "banking_app/Admin"
	"fmt"
)

func main() {
	Admin := admin.Newadmin("Vishav", "Pathaia")
	fmt.Println(Admin)
	axisBank := Admin.CreateNewBank("Axis_Bank")
	fmt.Println(axisBank)
	Aniket := Admin.CreateNewCustomer("Aniket", "Pardeshi", 1)
	fmt.Println(Aniket)
	iciciBank := Admin.CreateNewBank("ICICI_Bank")
	fmt.Println(iciciBank)

	AniketAccount2, aerr := Aniket.AddNewAccount(iciciBank)
	if aerr != nil {
		fmt.Println(aerr)
	}
	fmt.Println(AniketAccount2)
	fmt.Println(Aniket.Accounts[1])
	fmt.Println("-------------------------------------")
	transactions := Admin.PassBook(Aniket.Customer_id, 1, 0)
	for _, val := range transactions {
		fmt.Println(val)
	}
	fmt.Println("-------------------------------------")
	err := Aniket.DepositMoney(5000, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Checking balance---->", Aniket.Accounts[0]) //id 1
	fmt.Println("Checking balance---->", Aniket.Accounts[1]) //id 2
	Aniket.TransferMoneyInternally(1, 2, 1000)
	fmt.Println("Checking balance---->", Aniket.Accounts[1])
	fmt.Println("____________________________")
	account1, _ := Aniket.GetAccountById(1)
	account2, _ := Aniket.GetAccountById(2)
	fmt.Println(account1.Bank.Fullname)
	fmt.Println(account2.Bank.Fullname)

	fmt.Println("<--------------------------------------------->")
	Someone := Admin.CreateNewCustomer("Some", "One", 1)
	fmt.Println(Someone)
	Admin.TransferMoney_To_External(1000, 1, 2, 2, 1)
	fmt.Println(Someone)
}
