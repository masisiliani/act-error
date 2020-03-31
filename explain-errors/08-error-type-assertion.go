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

type dontKnowError struct{ err error }

func (myErr dontKnowError) Error() string {
	return "i don't know what's happening"
}

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		err := &myError{functionName: "divide", err: errors.New("Can not divide by Zero")}
		return 0, err
	}
	return (number1 / number2), nil
}

func processReturnUnknowError() error {
	return &dontKnowError{errors.New("Errors test")}
}

func main() {

	err := processReturnUnknowError()

	//// qualquer outro processo que retornasse um erro desconhecido
	if _, ok := err.(*dontKnowError); ok {
		fmt.Println(err)
	}

	var number1 int = 4
	var number2 int = 0

	result, err := divide(number1, number2)

	switch err.(type) {
	case *myError:
		fmt.Println("myError: ", err.Error())
	case *dontKnowError:
		fmt.Println("dontKnowError: ", err.Error())
	default:
		fmt.Println("%d/%d is %d", number1, number2, result)
	}

}
