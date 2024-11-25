package main

import (
	"fmt"
	"go-alura-curso-oo/contas"
)

func main() {
	contaDoRodrigo := contas.ContaCorrente{
		Titular:       "Rodrigo",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         300,
	}

	contaDaIngryd := contas.ContaCorrente{
		Titular:       "Ingryd",
		NumeroAgencia: 589,
		NumeroConta:   654321,
		Saldo:         100,
	}

	fmt.Println(contaDoRodrigo)
	fmt.Println(contaDaIngryd)

	status := contaDoRodrigo.Transferir(100, &contaDaIngryd)
	fmt.Println(status)

	fmt.Println(contaDoRodrigo)
	fmt.Println(contaDaIngryd)
}
