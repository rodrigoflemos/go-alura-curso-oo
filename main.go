package main

import (
	"fmt"
	"go-alura-curso-oo/clientes"
	"go-alura-curso-oo/contas"
)

func main() {

	clienteRodrigo := clientes.Titular{
		Nome:      "Rodrigo",
		CPF:       "000.000.000-00",
		Profissao: "Desenvolvedor",
	}

	contaDoRodrigo := contas.ContaCorrente{
		Titular:       clienteRodrigo,
		NumeroAgencia: 456,
		NumeroConta:   123456,
		Saldo:         1000,
	}

	fmt.Println(contaDoRodrigo)
}
