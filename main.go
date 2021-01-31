package main

import "fmt"

// ContaCorrente ...
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// ContaLucas := ContaCorrente{titular: "Lucas", numeroConta: 123456, saldo: 125.5}
	ContaLucas := ContaCorrente{"Lucas", 589, 123456, 125.5}

	fmt.Println(ContaLucas)

	// "*" repressenta o ponteiro
	var ContaDaCris *ContaCorrente
	ContaDaCris = new(ContaCorrente)

	ContaDaCris.titular = "Cris"

	var ContaDaCris2 *ContaCorrente
	ContaDaCris2 = new(ContaCorrente)

	ContaDaCris2.titular = "Cris"

	fmt.Println(&ContaDaCris)
	fmt.Println(&ContaDaCris2)

	// true devido ao uso do ponterio, é efetuado a verificação do conteúdo.
	fmt.Println(*ContaDaCris == *ContaDaCris2)

	// false porque os endereços de memória são diferentes
	fmt.Println(ContaDaCris == ContaDaCris2)

}
