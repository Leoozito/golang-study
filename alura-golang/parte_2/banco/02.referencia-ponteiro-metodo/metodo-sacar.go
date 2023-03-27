package main

import (
	"fmt"
	"alura-golang/parte_2/contas"
)

func main() {
	contaLeozito := ContaCorrente{titular: "Leonardo Nogueira", saldo: 500000.00}
	contaVitoria := ContaCorrente{titular: "Vitoria Regia", saldo: 60000.00}

	status := contaLeozito.Transferir(200, &contaVitoria)

	fmt.Println(status)
	fmt.Println(contaLeozito)
	fmt.Println(contaVitoria)
	// contaLeozito.titular = "Leonardo Nogueira"
	// contaLeozito.saldo = 2000.00

	// fmt.Println(contaLeozito.saldo)

	// fmt.Println(contaLeozito.Sacar(quantoAsacar))
	// fmt.Println(contaLeozito.saldo)

	// fmt.Println(contaLeozito.Depositar(quantoAdepositar))
	// fmt.Println(contaLeozito.saldo)
}



