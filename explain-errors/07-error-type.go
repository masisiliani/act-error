package main

import (
	"errors"
	"fmt"
)

type myError struct {
	functionName string
	err          error
}

func (myErr myError) Error() string {
	return fmt.Sprintf("Error: %s | function: %s", myErr.err.Error(), myErr.functionName)
}

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		err := myError{functionName: "divide", err: errors.New("Can not divide by Zero")}
		return 0, err
	}
	return (number1 / number2), nil
}

func main() {

	var number1 int = 4
	var number2 int = 0

	result, err := divide(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d/%d is %d", number1, number2, result)
	}
}
