package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRodrigo := ContaCorrente{
		titular:       "Rodrigo",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

	contaDaMaria := ContaCorrente{
		"Maria", 589, 123456, 125.50,
	}

	//Outra maneira de inicializar struct ( pouco utilizado )
	var contaDaJoana *ContaCorrente
	contaDaJoana = new(ContaCorrente)
	contaDaJoana.titular = "Joana"

	fmt.Println(contaDoRodrigo)
	fmt.Println(contaDaMaria)
	fmt.Println(contaDaJoana)
}
