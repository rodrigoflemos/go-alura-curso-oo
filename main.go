package main

import (
	"fmt"
	"go-alura-curso-oo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoRodrigo := contas.ContaPoupanca{}

	contaDoRodrigo.Depositar(100)

	PagarBoleto(&contaDoRodrigo, 60)

	fmt.Println(contaDoRodrigo.ObterSaldo())

	contaDaMaria := contas.ContaCorrente{}

	contaDaMaria.Depositar(100)

	PagarBoleto(&contaDaMaria, 120)

	fmt.Println(contaDaMaria.ObterSaldo())
}
