package main

import (
	"errors"
	"fmt"
)

func main() {

	// criando o erro
	// myErr := errors.New("Algo de errado não está certo.")
	myErr := errors.New("algo de errado não está certo")

	// escrever mensagem de erro
	fmt.Println(myErr)
}
