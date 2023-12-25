package interview_cases

import (
	"fmt"
)

func myFunc() error {
	var err *MyError
	if !someCondition() {
		err = &MyError{"an error occurred"}
	}
	return err
}

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func someCondition() bool {
	return false
}

func GetError() {
	err := myFunc()
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error is not nil")
	} else {
		fmt.Println("Error is nil")
	}
}
