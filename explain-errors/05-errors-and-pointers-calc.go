package main

import (
	"errors"
	"fmt"
)

func divide(number1 int, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("Can not divide number by Zero")
	}
	return (number1 / number2), nil
}

func main() {

	var number1 int = 4
	var number2 int = 0

	_, err1 := divide(number1, number2)
	_, err2 := divide(number1, number2)

	fmt.Println(err1 == err2)

	// fmt.Println("err1", &err1)
	// fmt.Println("err2", &err2)

	// fmt.Println(err1.Error() == err2.Error())

	// fmt.Println("err1", err1)
	// fmt.Println("err2", err2)

}
