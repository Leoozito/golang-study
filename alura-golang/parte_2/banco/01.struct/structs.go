								// BASICO

// ajudar uma agencia de banco gerenciar as contas

package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numDaConta int
	saldo float64
}

func main() {
	contaLeozito := ContaCorrente{titular: "Leonardo Nogueira", numeroAgencia: 589, numDaConta: 123456, saldo: 20785050.50 }
	contaVitoriaRegia := ContaCorrente{"Vitoria Regia", 339, 125739, 1014578.78 }
	
	fmt.Println(contaLeozito)
	fmt.Println(contaVitoriaRegia)
}