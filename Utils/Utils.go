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

func GenerateUniqueID() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(90000) + 10000
}
