package admin

import (
	bank "banking_app/Bank"
	customer "banking_app/Customer"
	utils "banking_app/Utils"
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
