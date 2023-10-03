package basics

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	privateField string
}

func GetPrivateFiled() {
	instance := MyStruct{"Hello, private field!"}
	valueOf := reflect.ValueOf(instance)
	field := valueOf.FieldByName("privateField")

	if field.IsValid() {
		if field.CanInterface() {
			fmt.Println("Accessed private field:", field.Interface())
		} else {
			fmt.Println("Cannot access private field.")
		}
	} else {
		fmt.Println("Field not found.")
	}
}
