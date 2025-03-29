package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	numeroConta int
	saldo float64
}

func ponteiros() {
	var conta *ContaCorrente
	conta = new(ContaCorrente)

	conta.titular = "Leonardo"
	conta.agencia = 5678
	conta.numeroConta = 912345
	conta.saldo = 3300.00

	fmt.Println(*conta)
}