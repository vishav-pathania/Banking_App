package utils

import (
	"fmt"
	"reflect"
)

func GetVariableType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic:", r)
	}
}
