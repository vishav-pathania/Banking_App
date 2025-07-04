package admin

import (
	utils "banking_app/Utils"
)

var adminMap = make(map[int]*Admin)
var adminMapid = -1

type Admin struct {
	Id        int
	FirstName string
	LastName  string
	IsActive  bool
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
	}
	adminMap[adminMapid] = ad
	return ad
}
