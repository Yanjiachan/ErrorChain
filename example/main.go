package main

import (
	ec "errorChain"
	"errors"
	"fmt"
)

func returnErr(*ec.ErrorChain) error {
	return errors.New("err")
}

func returnNil(*ec.ErrorChain) error {
	return nil
}

func otherCheck(i int) func(*ec.ErrorChain) error {
	return func(*ec.ErrorChain) error {
		if i < 0 {
			return errors.New("i < 0")
		}
		return nil
	}
}

func main() {
	e := &ec.ErrorChain{}
	err := e.CheckError(otherCheck(1), returnNil, returnErr).Error
	fmt.Println(err)
}
