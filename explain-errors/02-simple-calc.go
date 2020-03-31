package main

import (
	"errors"
	"fmt"
)

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("Can not divide by Zero")
	}
	return (number1 / number2), nil
}

func main() {

	result, err := divide(4, 2)

	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("4/2 is", result)
	}
}
