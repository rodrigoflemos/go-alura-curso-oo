package contas

import "go-alura-curso-oo/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possivel realizar o saque"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {

	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor do que zero", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
