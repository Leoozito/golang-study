package main

import (
	"fmt"
)

func main () {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa !")
	} else {
		fmt.Println("Comando invalido")
	}
}