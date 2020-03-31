package main

import (
	"errors"
	"fmt"
)

func main() {

	err1 := errors.New("Erro 1")
	err2 := fmt.Errorf("Erro 2: %v", err1)
	err3 := fmt.Errorf("Erro 3: %w", err1)

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)

	err4 := errors.Unwrap(err2)
	err5 := errors.Unwrap(err3)

	fmt.Println(err4)
	fmt.Println(err5)

}
