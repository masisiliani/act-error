package main

import (
	"errors"
	"fmt"
)

func main() {

	err1 := errors.New("Erro 1")
	err2 := fmt.Errorf("Erro 2: %w", err1)
	err3 := fmt.Errorf("Erro 3: %w", err2)
	err4 := fmt.Errorf("Erro 4: %w", err3)
	err5 := fmt.Errorf("Erro 5: %w", err4)

	err6 := errors.Unwrap(err5)
	fmt.Println(err6)

}
