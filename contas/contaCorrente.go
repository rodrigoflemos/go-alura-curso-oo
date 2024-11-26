package contas

import "go-alura-curso-oo/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possivel realizar o saque"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor do que zero", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
