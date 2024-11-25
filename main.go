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

	contaDoRodrigo2 := ContaCorrente{
		titular:       "Rodrigo",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

	// == compara conteudo em Go
	fmt.Println(contaDoRodrigo == contaDoRodrigo2)

	//Outra maneira de inicializar struct
	var contaDaJoana *ContaCorrente
	contaDaJoana = new(ContaCorrente)
	contaDaJoana.titular = "Joana"

	var contaDaJoana2 *ContaCorrente
	contaDaJoana2 = new(ContaCorrente)
	contaDaJoana2.titular = "Joana"

	fmt.Println(contaDaJoana)
	fmt.Println(contaDaJoana2)
	fmt.Println(&contaDaJoana)
	fmt.Println(&contaDaJoana2)

	//A comparacao é realizada por endereco de memoria = false
	fmt.Println(contaDaJoana == contaDaJoana2)
	//A comparacao é realizada por conteudo = true
	fmt.Println(*contaDaJoana == *contaDaJoana2)

}
