package main

import (
	customer "banking_app/Customer"
	"fmt"
)

func main() {
	newAdmin := customer.Newadmin("Vishav", "Pathania")
	fmt.Println(newAdmin)
	AxisBank := newAdmin.CreateNewBank("Axis_Bank")
	ICICIBank := newAdmin.CreateNewBank("ICICI_Bank")
	fmt.Println(AxisBank)
	fmt.Println(ICICIBank)
	Aniket := newAdmin.CreateNewCustomer("Aniket", "Pardeshi", AxisBank.Bank_id)
	fmt.Println(Aniket)
	AniketAccount2 := Aniket.AddNewAccount(ICICIBank.Bank_id)
	fmt.Println(AniketAccount2)
	fmt.Println(Aniket)
	// Aniket.DeleteAccountById(AniketAccount2.Account_No)
	// fmt.Println(Aniket)
	Aniket.DepositMoney(5000, AniketAccount2.Account_No)
	fmt.Println(Aniket)
	Aniket.WithDrawMoney(500, AniketAccount2.Account_No)
	fmt.Println(Aniket)
	fmt.Println("-----------------------")
	fmt.Println(AniketAccount2)
	Aniket.TransferMoneyInternally(AniketAccount2.Account_No, Aniket.Accounts[0].Account_No, 600)
	fmt.Println(AniketAccount2)
	fmt.Println(Aniket.Accounts[0])
	// SomeOneNew := newAdmin.CreateNewCustomer("SomeOne", "New", ICICIBank.Bank_id)
	SomeOneNew := newAdmin.CreateNewCustomer("SomeOne", "New", AxisBank.Bank_id)
	fmt.Println(SomeOneNew)
	Aniket.TransferMoney_To_External(600, Aniket.Customer_id, SomeOneNew.Customer_id, AniketAccount2.Account_No, SomeOneNew.Accounts[0].Account_No)
	fmt.Println(Aniket)
	fmt.Println(SomeOneNew)
	// newAdmin.DeleteBank(ICICIBank.Bank_id)
	SomeOneNew.DeleteAccountById(SomeOneNew.Accounts[0].Account_No)
	newAdmin.DeleteBank(ICICIBank.Bank_id)
	fmt.Println(ICICIBank) //bank status false now
	resultpage := newAdmin.GetPassBook_ById(Aniket.Customer_id, AniketAccount2.Account_No, 0)
	fmt.Println("passbook or transactions below--------------------------------------")
	for _, vals := range resultpage {
		fmt.Println(vals)
	}
	fmt.Println("for ledger--------------------->")
	AxisBankLedgerRecords := newAdmin.GetLedgerByBank_Id(ICICIBank.Bank_id, 0)
	for _, ledgerVals := range AxisBankLedgerRecords {
		fmt.Println(ledgerVals)
	}
	amount := newAdmin.SettleMent(AxisBank.Bank_id, ICICIBank.Bank_id)
	fmt.Println("settlement amount--------------->", amount)
}
