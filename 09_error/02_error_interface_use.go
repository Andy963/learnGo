package main

import (
	"errors"
	"fmt"
)

func myDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("b can not be zero")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := myDiv(10, 0)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("result=", result)
	}
}
