package main

import "fmt"

// ContaCorrente ...
type ContaCorrente struct {
	titular          string
	TitularCPF       string
	TitularProfissao string
	numeroAgencia    int
	numeroConta      int
	saldo            float64
}

func main() {

	ContaDoLucas := ContaCorrente{}
	ContaDoLucas.titular = "Lucas"
	ContaDoLucas.saldo = 200

	fmt.Println(ContaDoLucas.saldo)

	fmt.Println(ContaDoLucas.Sacar(100))

	status, valor := ContaDoLucas.Depositar(100)

	fmt.Println(status, valor)

}

// Sacar ...
func (c *ContaCorrente) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "Saldo insuficiente!!"
	}
}

// Depositar ...
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito efetuado com sucesso!!", c.saldo
	} else {
		return "Não foi possível efetuar o depósito!!", c.saldo
	}
}
