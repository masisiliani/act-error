package main

import (
	"fmt"
)

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		return 0, fmt.Errorf("Can not divide number %d by Zero", number1)
	}
	return (number1 / number2), nil
}

func main() {

	var number1 int = 4
	var number2 int = 0

	result, err := divide(number1, number2)

	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Printf("%d/%d is %d", number1, number2, result)
	}
}
