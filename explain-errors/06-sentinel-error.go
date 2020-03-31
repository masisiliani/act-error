package main

import (
	"errors"
	"fmt"
)

var errorCanNotDivide = errors.New("Can not divide by Zero")

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errorCanNotDivide
	}
	return (number1 / number2), nil
}

func main() {

	var number1 int = 4
	var number2 int = 0

	_, err1 := divide(number1, number2)

	fmt.Println(err1 == errorCanNotDivide)

}
