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

func main() {
	contaDoRodrigo := ContaCorrente{
		titular:       "Rodrigo",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         500,
	}
	fmt.Println(contaDoRodrigo.saldo)
	fmt.Println(contaDoRodrigo.Sacar(-100))
	fmt.Println(contaDoRodrigo.saldo)
}
