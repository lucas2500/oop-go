package main

import (
	contas "banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valorBoleto float64) string
}

func main() {

	//-- Conta Corrente

	// ContaDoLucas := c.ContaCorrente{Titular: "Lucas", Saldo: 300}
	// ContaDoRafa := c.ContaCorrente{Titular: "Rafael", Saldo: 400}

	// status := ContaDoLucas.Transferir(100.00, &ContaDoRafa)

	// fmt.Println(status)
	// fmt.Println(ContaDoLucas)
	// fmt.Println(ContaDoRafa)

	// ClienteLucas := cliente.Titular{
	// 	Nome:      "Lucas",
	// 	CPF:       "11790143401",
	// 	Profissao: "Programador",
	// }
	// ContaDoLucas := contas.ContaCorrente{ClienteLucas, 123, 123456, 100}

	// ContaDoLucas := contas.ContaCorrente{}

	// ContaDoLucas.Depositar(500)

	// ContaDoLucas.Sacar(200)

	// fmt.Println(ContaDoLucas.ConsultarSaldo())

	//-- Conta poupança

	// ContaDoLucas := contas.ContaPoupanca{}

	// ContaDoLucas.Depositar(200.00)
	// fmt.Println(ContaDoLucas.ConsultarSaldo())

	ContaDoLucas := contas.ContaPoupanca{}
	ContaDoLucas.Depositar(500)

	PagarBoleto(&ContaDoLucas, 60)

	fmt.Println(ContaDoLucas.ConsultarSaldo())

	ContaDoLucas2 := contas.ContaPoupanca{}

	ContaDoLucas2.Depositar(600)

	PagarBoleto(&ContaDoLucas2, 300)

	fmt.Println(ContaDoLucas2.ConsultarSaldo())

}

// PagarBoleto ...
func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
