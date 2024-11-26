package main

import (
	"fmt"
	"go-alura-curso-oo/contas"
)

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo)
	fmt.Println(contaExemplo.ObterSaldo())
}
