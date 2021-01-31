package contas

import (
	"banco/clientes"
)

// ContaCorrente ...
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
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

// Transferir ...
func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {

	if valor < c.saldo && valor > 0 {

		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true

	} else {

		return false
	}
}

// ConsultarSaldo ...
func (c *ContaCorrente) ConsultarSaldo() float64 {

	return c.saldo
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
