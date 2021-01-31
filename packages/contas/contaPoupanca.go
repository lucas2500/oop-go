package contas

import "banco/clientes"

// ContaPoupanca ...
type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

// Sacar ...
func (c *ContaPoupanca) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!!"
	} else {
		return "Saldo insuficiente!!"
	}
}

// Depositar ...
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito efetuado com sucesso!!", c.saldo
	} else {
		return "Não foi possível efetuar o depósito!!", c.saldo
	}
}

// ConsultarSaldo ...
func (c *ContaPoupanca) ConsultarSaldo() float64 {

	return c.saldo
}
