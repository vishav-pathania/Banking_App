package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func GetVariableType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic:", r)
	}
}

func GenerateTransactionID() int64 {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Int63n(9000000000) + 1000000000 // 10-digit number
}
