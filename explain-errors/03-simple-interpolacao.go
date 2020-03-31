package main

import (
	"fmt"
)

func main() {

	// criando o erro
	myErr := showPersonalError("Silvana", 123)

	// escrever mensagem de erro
	fmt.Println(myErr)

}

func showPersonalError(name string, logNumber int) error {
	return fmt.Errorf("%s, algo de errado não está certo. Verifique o log %d", name, logNumber)
}
