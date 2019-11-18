package main

import (
	"errors"
	"fmt"
)

var errDidNotWork = errors.New("did not work")

// DoTheThing ...
func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = errDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", errDidNotWork
}

func scopetest() {
	// 因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量
	fmt.Println(DoTheThing(true))  // nil
	fmt.Println(DoTheThing(false)) // nil
}
