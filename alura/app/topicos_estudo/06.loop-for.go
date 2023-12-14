package main

import (
	"fmt"
	"os"
)

func main() {
	//  no GO não tem "WHILE", ou seja, usando o for pode se fazer loop infinito
	for {
		exibeMenu()

		comando := retornaComando()

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa !")
			os.Exit(0)
		default:
			fmt.Println("Comando invalido")
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func retornaComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
