package main

import "fmt"

func main() {
	noStruct()
	yesStruct()
}

func noStruct() {
	titular := "Leonardo"
	agencia := 1234
	numeroConta := 12345678
	saldo := 2300.00

	fmt.Println(titular, agencia, numeroConta, saldo)

	titular2 := "Vitoria"
	agencia2 := 5678
	numeroConta2 := 912345
	saldo2 := 3300.00

	fmt.Println(titular2, agencia2, numeroConta2, saldo2)
}

type ContaCorrente struct {
	titular string
	agencia int
	numeroConta int
	saldo float64
}

func yesStruct() {
	conta := ContaCorrente{
		"Leonardo",
		1234,
		12345678,
		2300.00,
	}

	conta2 := ContaCorrente{
		"Vitoria",
		5678,
		912345,
		3300.00,
	}

	fmt.Println(conta)
	fmt.Println(conta2)
}