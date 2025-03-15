// Este arquivo mostra: Sobre o "print" e "scan" (que armazena dados que são informados)

package main

import "fmt"

func main() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando) // o "&" é referênciando o endereço da variavel
	// O "Scan" serve para capitar o que digitamos.

	// O que digitarmos será armazenado no endereço da variavel escolhida
	fmt.Println("O comando escolhido foi: ", comando)
}
