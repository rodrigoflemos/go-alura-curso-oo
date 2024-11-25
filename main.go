package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	if valorDoSaque < 0 {
		return "O valor do saque precisa ser positivo"
	}
	podeSacar := valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor do que zero", c.saldo
	}

}

func main() {
	contaDoRodrigo := ContaCorrente{
		titular:       "Rodrigo",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         500,
	}
	fmt.Println(contaDoRodrigo.saldo)
	fmt.Println(contaDoRodrigo.Sacar(300))
	fmt.Println(contaDoRodrigo.saldo)
	fmt.Println(contaDoRodrigo.Depositar(800))
	status, valor := contaDoRodrigo.Depositar(500)
	fmt.Println(status, valor)
}
